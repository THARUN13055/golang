package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Connect to MongoDB
func connectDB() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb+srv://tharun:tharun@golang.vpv6p.mongodb.net/?retryWrites=true&w=majority&appName=golang")
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
}

// Get all collections from the specified database
func getAllCollections(c *gin.Context) {
	dbName := "sample_mflix" // Replace with your database name
	collections, err := client.Database(dbName).ListCollectionNames(context.TODO(), bson.M{})
	for _, a := range collections {
		if a == "movies" {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "working file founded"})
		}
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"collections": collections})
}

func main() {
	// Connect to MongoDB
	connectDB()

	// Set up Gin router
	router := gin.Default()

	// Define routes
	router.GET("/collections", getAllCollections)

	// Start the server
	router.Run(":6000")
}
