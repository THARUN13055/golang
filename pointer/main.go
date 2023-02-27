package main

import "fmt"

func main() {
	var number = 4
	var ptr = &number
	fmt.Println("enter the pointer", ptr)
	fmt.Println("enter the pointer", *ptr)
	fmt.Println("the number is ", number)
}
