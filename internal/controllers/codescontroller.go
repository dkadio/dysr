package controllers

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"

	//"io/ioutil"
	"log"
	//"net/http"

	classes "github.com/dkadio/dysr/internal"
	"github.com/dkadio/dysr/util"

	dm "github.com/dkadio/dysr/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionName = "codes"

type DataBaseAdapter interface {
	GetValueFor(string) string
	PutValueFor(string, string) error
}

type CodesController struct {
	adapter        DataBaseAdapter
	collection     *mongo.Collection
	collectionName string
}

func NewCodesController() CodesController {
	bolt := classes.NewBoltAdapterByString("redirects.db", "codes")
	return CodesController{adapter: bolt}
}

func NewMongoCodesController() CodesController {
	collection := util.GetCollection(util.GetDatabase(), collectionName)
	return CodesController{collection: collection, collectionName: collectionName}

}

func (c CodesController) CreateCode(gc *gin.Context, params *dm.CreateCode) (*dm.UserCode, error) {
	jwtclaims, _ := gc.Get("claims")
	claims := jwtclaims.(jwt.MapClaims)

	username := claims["preferred_username"].(string)

	code := dm.NewUserCode(username, params.Value, params.Options)

	//Put value to redirect store
	//c.adapter.PutValueFor(code.Code.Key, code.Code.Value)
	c.collection.InsertOne(context.TODO(), code)
	//TODO Publish Event for new key value
	return &code, nil
}

func (c CodesController) UpdateCode(gc *gin.Context, params *dm.Code) (*dm.UserCode, error) {
	filter := bson.M{"_id": bson.M{"$eq": params.UUID}}
	update := bson.M{"$set": bson.M{"value": params.Value, "options": params.Options, "updated": time.Now().Unix()}}
	log.Printf("Call with parms: %+v", params)

	result := dm.UserCode{}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	singleresult := c.collection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&result)
	log.Println("singleresult", singleresult)

	log.Println("Call", result)
	//id := gc.Param("id")
	//fmt.Printf("params: %v", params)
	return &result, nil
}

//Query all codes for user request
func (c CodesController) GetCodes(gc *gin.Context) ([]dm.UserCode, error) {

	result := []dm.UserCode{}
	filter := bson.D{{}}
	codes, err := c.collection.Find(context.TODO(), filter)
	if err != nil {
		return result, err
	}

	for codes.Next(context.TODO()) {
		uc := dm.UserCode{}
		codes.Decode(&uc)
		result = append(result, uc)
	}

	codes.Close(context.TODO())

	return result, nil
}

func (c CodesController) GetCode(gc *gin.Context, params *dm.Code) (*dm.UserCode, error) {
	result := dm.UserCode{}
	filter := bson.M{"_id": bson.M{"$eq": params.UUID}}
	c.collection.FindOne(context.TODO(), filter).Decode(&result)
	return &result, nil
}

func (c CodesController) DeleteCode(gc *gin.Context, params *dm.UserCode) error {
	log.Println("Call")
	filter := bson.M{"_id": bson.M{"$eq": params.UUID}}
	result, err := c.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	log.Println("Delete Count: ", result)
	return nil
}
