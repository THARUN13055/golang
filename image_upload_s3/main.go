package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env is not loading")
		return
	}
}

func main() {
	// fileopen, err := os.Open("./moon.jpg")
	// if err != nil {
	// 	log.Fatalf("The file is not opening: %v", err)
	// }

	// defer fileopen.Close()

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		log.Fatalf("region is not loading: %v", err)
	}

	svc := s3.NewFromConfig(cfg)

	// input := &s3.PutObjectInput{
	// 	Bucket:      aws.String(os.Getenv("S3_BUCKET")),
	// 	Key:         aws.String("/app/moon.jpg"),
	// 	ContentType: aws.String("image/jpeg"),
	// 	Body:        fileopen,
	// }

	// _, err = svc.PutObject(context.TODO(), input)

	// if err != nil {
	// 	log.Fatalf("The PutObject is not working: %v", err)
	// }
	// fmt.Println("s3 images has been uploaded successfully")
	mongodbclient, err := connectMongos()
	if err != nil {
		log.Fatalf("Error while connection of mongodb: %v", err)
	}

	uploadfile(svc, mongodbclient, "./moon.jpg", "user", "moon.jpg", "tharun")

}

func connectMongos() (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI("mongodb://tharun:password12345@3.7.247.50:27017/tharun").SetServerAPIOptions(serverAPI)
	// create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// sending the ping that it is working or not

	if err := client.Database("tharun").RunCommand(context.TODO(), bson.M{"ping": 1}).Err(); err != nil {
		panic(err)
	}

	fmt.Println("Database is successfully connected")

	return client, err

}

func uploadfile(svc *s3.Client, client *mongo.Client, filepath string, s3filepath string, filename string, author string) {

	fileopen, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("The file is not opening: %v", err)
	}

	defer fileopen.Close()

	uniquiddb := uuid.New().String()

	s3Key := fmt.Sprintf("%s/%s", s3filepath, filename)

	input := &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("S3_BUCKET")),
		Key:         aws.String(s3Key),
		ContentType: aws.String("image/jpeg"),
		Body:        fileopen,
	}

	_, err = svc.PutObject(context.TODO(), input)

	if err != nil {
		log.Fatalf("The PutObject is not working: %v", err)
	}
	fmt.Println("s3 images has been uploaded successfully")

	// store metadata to mongodb collection
	s3urlGenerate := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s/%s", os.Getenv("S3_BUCKET"), os.Getenv("AWS_REGION"), s3filepath, filename)

	collection := client.Database("tharun").Collection("images")
	_, err = collection.InsertOne(context.TODO(), bson.M{
		"id":     uniquiddb,
		"path":   s3Key,
		"s3url":  s3urlGenerate,
		"author": author,
	})
	if err != nil {
		log.Fatalf("The data is not inserted in db: %v", err)
	}
	fmt.Println("Images metadata is stored in mongodb successfully")
}

