package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Jsonfile struct {
	Tharun   string `json:"tharun"`
	Location string `json:"location"`
	Number   int    `json:"number"`
	Company  string `json:"company"`
	Role     string `json:"role"`
}

func main() {
	file, err := os.Open("result.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decode := json.NewDecoder(file)
	var data []Jsonfile
	if err := decode.Decode(&data); err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(reflect.TypeOf(data))
}
