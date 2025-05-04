package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type person struct {
		Name string
		Age  int
		Car  []string
	}

	// Create a person instance and marshal it to JSON
	p1 := &person{Name: "tharun", Age: 23, Car: []string{"tata", "audi"}}
	data, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(data))

	// Unmarshal the JSON data back into a new person instance
	p2 := &person{}
	err = json.Unmarshal(data, p2) // Use 'data' instead of 'data1'
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Println(p2.Name)
}
