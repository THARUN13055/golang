package main

import (
	"encoding/xml"
	"fmt"
)

type movie struct {
	Name  string
	Class string
}

func main() {
	p := &movie{"tharun", "12th"}
	value, _ := xml.MarshalIndent(p, "", " ")
	fmt.Println(string(value))
}
