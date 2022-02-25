package controllers

import (
	"context"
	//"github.com/golang-jwt/jwt/v4"
	//"time"

	//"io/ioutil"
	"log"
	//"net/http"

	"github.com/dkadio/dysr/util"

	"github.com/dkadio/dysr/internal/models"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type StatsController struct {
	collection *mongo.Collection
	nats       *nats.EncodedConn
}

//func NewStatsController() StatsController {
//	bolt := classes.NewBoltAdapterByString("redirects.db", "codes")
//	return StatsController{adapter: bolt}
//}

func NewStatsController() StatsController {
	collection := util.GetCollection(util.GetDatabase(), util.STATS_COLLECTION_NAME)
	nats, _ := util.GetNatsClient()
	return StatsController{nats: nats, collection: collection}
}

func (c StatsController) RegisterForEvents() {
	log.Println("Registering for Events....")
	c.nats.Subscribe(util.REDIRECTED_EVENT, c.handleRedirectEvent)
}

func (c StatsController) handleRedirectEvent(msg *models.Request) {
	//TODO handle the requstt
	c.updateStats(msg)
}

func (c StatsController) updateStats(msg *models.Request) {

	date := time.Now().Format("2006-01-02")

	filter := bson.M{"key": bson.M{"$eq": msg.CODE}, "date": bson.M{"$eq": date}}
	update := bson.M{"$inc": bson.M{"stats.clicks.total": 1}}

	result := models.Stats{}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	c.collection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&result)
	log.Println(result)

}

//func (c StatsController) getCodeId(key string) (string, error) {
//	result := models.UserCode{}
//	filter := bson.M{"key": bson.M{"$eq": key}}
//	c.collection.FindOne(context.TODO(), filter).Decode(&result)
//	log.Println("Found
//	return result., nil
//}
