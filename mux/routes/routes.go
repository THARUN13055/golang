package routes

import (
	"github.com/gorilla/mux"
	"github.com/tharun13055/golang/mux/handler"
	"github.com/tharun13055/golang/mux/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(r *mux.Router, db *mongo.Database) {

	// create the service
	userService := services.UserService(db)

	// create the handler

	userHandler := handler.UserHandler(userService)

	r.HandleFunc("/Health", handler.Health).Methods("GET").Host("localhost:5000").Schemes("http")
	r.HandleFunc("/GetUser", userHandler.GetUser).Methods("GET").Host("localhost:5000").Schemes("http")
}
