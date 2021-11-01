package controllers

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"

	//"io/ioutil"
	"log"
	//"net/http"

	"github.com/dkadio/dysr/util"

	dm "github.com/dkadio/dysr/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CodesController struct {
	collection *mongo.Collection
	nats       *nats.EncodedConn
}

//func NewCodesController() CodesController {
//	bolt := classes.NewBoltAdapterByString("redirects.db", "codes")
//	return CodesController{adapter: bolt}
//}

func NewMongoCodesController() CodesController {
	collection := util.GetCollection(util.GetDatabase(), util.CODES_COLLECTION_NAME)
	nats, _ := util.GetNatsClient()
	return CodesController{nats: nats, collection: collection}

}

func (c CodesController) CreateCode(gc *gin.Context, params *dm.CreateCode) (*dm.UserCode, error) {
	jwtclaims, _ := gc.Get("claims")
	claims := jwtclaims.(jwt.MapClaims)

	username := claims["preferred_username"].(string)

	code := dm.NewUserCode(username, params.Value, params.Options)

	//Put value to redirect store
	//c.adapter.PutValueFor(code.Code.Key, code.Code.Value)
	c.collection.InsertOne(context.TODO(), code)
	msg := &dm.CodeValue{Key: code.Code.Key, Value: code.Code.Value}
	c.nats.Publish(util.CODES_CREATED_EVENT, &msg)
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

	c.collection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&result)

	log.Println("Call", result)
	//id := gc.Param("id")
	//fmt.Printf("params: %v", params)
	msg := &dm.CodeValue{Key: result.Code.Key, Value: result.Code.Value}
	c.nats.Publish(util.CODES_UPDATED_EVENT, msg)
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

	oldcode, _ := c.GetCode(nil, &dm.Code{UUID: params.UUID})
	result, err := c.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	log.Println("Delete Count: ", result)
	c.nats.Publish(util.CODES_DELETE_EVENT, oldcode.Code.Key)
	return nil
}
