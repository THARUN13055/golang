package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Env is not loading")
	}

	db := os.Getenv("DATABASE")
	fmt.Println(db)

}
