package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectDb() *mongo.Client {

	Client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	} else {
		MongoClient = Client
		fmt.Println("Connected to DataBase")
		return Client
	}

}

func ConnectToCollection(collectionName string) *mongo.Collection {
	Client := ConnectDb()
	MongoDataBase := Client.Database("EduCoApp")
	collection := MongoDataBase.Collection(collectionName)
	return collection
}
