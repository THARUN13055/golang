package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env is not loading")
		return
	}
}

func main() {
	fileopen, err := os.Open("./moon.jpg")
	if err != nil {
		log.Fatalf("The file is not opening: %v", err)
	}

	defer fileopen.Close()

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		log.Fatalf("region is not loading: %v", err)
	}

	svc := s3.NewFromConfig(cfg)

	input := &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("S3_BUCKET")),
		Key:         aws.String("/app/moon.jpg"),
		ContentType: aws.String("image/jpeg"),
		Body:        fileopen,
	}

	_, err = svc.PutObject(context.TODO(), input)

	if err != nil {
		log.Fatalf("The PutObject is not working: %v", err)
	}
	fmt.Println("s3 images has been uploaded successfully")
}
