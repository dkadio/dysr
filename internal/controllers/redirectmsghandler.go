package controllers

import (
	"context"
	"errors"
	"github.com/dkadio/dysr/internal"
	"github.com/dkadio/dysr/internal/models"
	"github.com/dkadio/dysr/util"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type DataBaseAdapter interface {
	GetValueFor(string) string
	PutValueFor(string, string) error
}

type RedirectController struct {
	adapter    DataBaseAdapter
	collection *mongo.Collection
	nats       *nats.EncodedConn
	config     util.Config
}

func NewRedirectController() RedirectController {
	bolt := classes.NewBoltAdapterByString("redirects.db", "codes")
	collection := util.GetCollection(util.GetDatabase(), util.CODES_COLLECTION_NAME)
	config := util.LoadConfig()
	con, _ := util.GetNatsClient()
	return RedirectController{nats: con, adapter: bolt, collection: collection, config: config}
}

func (r RedirectController) subscribeCodeEvents() {
	r.nats.Subscribe(util.CODES_CREATED_EVENT, r.handleCreate)
	r.nats.Subscribe(util.CODES_UPDATED_EVENT, r.handleUpdate)
	r.nats.Subscribe(util.CODES_DELETE_EVENT, r.handleDelete)
}

func (r RedirectController) GetValueFor(key string) (string, error) {
	log.Println("Searching for: ", key)
	v := r.adapter.GetValueFor(key)
	if v != "" {
		return v, nil
	}
	log.Println("Return:", r.config.FrontEndUrl)
	//TODO Notification to user that no Redirect is set
	return r.config.FrontEndUrl, errors.New("Code not fount")
}

func (r RedirectController) InformRedirect(req models.Request) {
	log.Println("Publish event,", req)
	err := r.nats.Publish(util.REDIRECTED_EVENT, req)
	if err != nil {
		log.Println(err)
	}
}

func (r RedirectController) loadIntialKeyStore() {
	filter := bson.D{{}}
	codes, err := r.collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Cant Load Data from mongo to store", err)
	}
	log.Println("loading into store: ", codes.RemainingBatchLength())

	for codes.Next(context.TODO()) {
		uc := models.UserCode{}
		codes.Decode(&uc)
		r.adapter.PutValueFor(uc.Code.Key, uc.Code.Value)
	}

	codes.Close(context.TODO())
}

func (r RedirectController) Init() {
	r.subscribeCodeEvents()
	r.loadIntialKeyStore()
}

func (r RedirectController) handleCreate(code *models.CodeValue) {
	r.adapter.PutValueFor(code.Key, code.Value)
	log.Printf("Received new Event %+v", code)
}

func (r RedirectController) handleUpdate(code *models.CodeValue) {
	r.adapter.PutValueFor(code.Key, code.Value)
	log.Printf("Received Update Event %+v", code)
}

func (r RedirectController) handleDelete(id string) {
	r.adapter.PutValueFor(id, "")
	log.Println("Delete on ID:", id)
}
