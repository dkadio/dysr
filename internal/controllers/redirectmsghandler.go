package controllers

import (
	"context"
	"github.com/dkadio/dysr/internal"
	"github.com/dkadio/dysr/internal/models"
	"github.com/dkadio/dysr/util"
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
}

func NewRedirectController() RedirectController {
	bolt := classes.NewBoltAdapterByString("redirects.db", "codes")
	collection := util.GetCollection(util.GetDatabase(), util.CODES_COLLECTION_NAME)
	return RedirectController{adapter: bolt, collection: collection}
}

func (r RedirectController) subscribeCodeEvents() {
	con, _ := util.GetNatsClient()
	con.Subscribe(util.CODES_CREATED_EVENT, r.handleCreate)
	con.Subscribe(util.CODES_UPDATED_EVENT, r.handleUpdate)
	con.Subscribe(util.CODES_DELETE_EVENT, r.handleDelete)
}

func (r RedirectController) GetValueFor(key string) string {
	return r.adapter.GetValueFor(key)
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
