package main

import (
	"fmt"
)

func main () {
	one := make(chan string)
	second := make(chan string)

	go goOne(one)
	go gotwo(second)

	select {
	case val1 := <- one:
		fmt.Println(val1)
	case val2 := <- second:
		fmt.Println(val2)
	default:
		fmt.Println("No values are running")
	}
}

func goOne(first chan string){
	first <- "channal values is 1"
}

func gotwo(second chan string){
	second <- "channal values is 2"
}