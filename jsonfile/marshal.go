package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	jsonData, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData)) // Output: {"name":"Alice","age":30}

	// we need to unmarshal

	err = json.Unmarshal(jsonData,&p)

	if err != nil {
		panic(err)
	}
	fmt.Print(p)

	fmt.Println()

}
