package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println(file.Name())

}
