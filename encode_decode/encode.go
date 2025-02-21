package main


import (
	"fmt"
	"encoding/json"
)


type data struct {
	Id int `json: "id"`
	Name string `json: "name"`
}

func main(){

	user := data {
		Id: 1,
		Name: "tharun",
	}

	// Here we need to marshal the data now
	// marchal means it will convert slice , list, map into json formate
	// Encode the data struct to JSON
	// Marshalling is the process of converting a Go data structure (like a struct, map, or slice) into a JSON format. In Go, this is done using the json.Marshal function from the encoding/json package.



	jsonData,err := json.Marshal(user)
	if err != nil {
		fmt.Println("The given data is not marshal",err)
		return
	}
	// Print the json data as []byte
	fmt.Println("the json data is :  ",jsonData)
	// Print the JSON data readable
	fmt.Println("the json data is :  ",string(jsonData))
}

// Encoding is a broader term that refers to the process of converting data into a specific format. In the context of JSON in Go, encoding typically refers to writing JSON data directly to an io.Writer, such as a file or an HTTP response.

// In Go, you can use the json.NewEncoder function to create a new JSON encoder that writes to a specified writer. This is useful when you want to write JSON data directly to a file or send it over a network.