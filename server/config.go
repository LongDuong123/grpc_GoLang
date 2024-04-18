package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var client *mongo.Client

func connectDatabase() error {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return err
	}
	collection = client.Database("BookSystem").Collection("Books")
	return nil
}
