package services

import (
	"context"
	"time"

	"github.com/tharun13055/golang/mux/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// interface is use for verify wether the function is there or not
type UserService interface {
	GetUserById(id string) (*models.User, error)
}

// It only contain the mongodb collection
type userService struct {
	Collection *mongo.Collection
}

func (s *userService) GetUserById(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User

	err := s.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
