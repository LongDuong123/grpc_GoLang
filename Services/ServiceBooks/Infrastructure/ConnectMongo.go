package infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func ConnectDatabase() (*Database, error) {
	clientMongo, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	err = clientMongo.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	collectionMongo := clientMongo.Database("BookSystem").Collection("Books")
	return &Database{collection: collectionMongo, client: clientMongo}, nil
}
