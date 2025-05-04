package configs

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func connectdb() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb+srv://tharun:tharun@golang.vpv6p.mongodb.net/?retryWrites=true&w=majority&appName=golang")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to mongoDB")

}
