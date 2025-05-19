package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tharun13055/golang/mux/db"
	"github.com/tharun13055/golang/mux/routes"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Server started at http://localhost:5000")

	// mongodb connection
	client := db.DB_Connection()
	if client == nil {
		log.Fatal("MongoDB connection failed")
	}
	// Seprate the route
	routes.RegisterRoutes(r)
	err := http.ListenAndServe("localhost:5000", r)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
