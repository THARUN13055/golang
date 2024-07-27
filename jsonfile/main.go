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
	p1 := &person{Name: "tharun", Age: 23, Car: []string{"tata", "audi"}}
	data,_:= json.Marshal(p1)
	fmt.Println(string(data))

	data1:={"Name":"tharun","Age":23,"Car":["tata","audi"]}

	p2 := &person{}
	json.Unmarshal([]byte(data1),p2)
	fmt.Println(p2.Name)

}
