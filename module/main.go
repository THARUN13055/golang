package main

import (
	"fmt"

	"example.com/main/sample"
	"example.com/main/sample/one"
)

func main() {
	fmt.Println(sample.Hello())
	fmt.Println(one.World())
}
