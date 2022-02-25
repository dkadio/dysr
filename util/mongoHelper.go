package util

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

const CODES_COLLECTION_NAME = "codes"
const STATS_COLLECTION_NAME = "stats"

//GetMongoClient - Return mongodb connection to work with
func GetMongoClient() (*mongo.Client, error) {
	//Perform connection creation operation only once.
	config := LoadConfig()
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(config.MongoUri)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

func GetDatabase() *mongo.Database {
	client, err := GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	config := LoadConfig()
	return client.Database(config.DatabaseName)
}

func GetCollection(db *mongo.Database, name string) *mongo.Collection {
	return db.Collection(name)
}
