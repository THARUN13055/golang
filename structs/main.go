package main

import "fmt"

func main() {
	fmt.Println("struct in golang")

	// go has no inheritance and super or parent

	tharun := User{"tharun", "tharun@we323.go", true, 32}
	fmt.Println(tharun)
	fmt.Printf("name is %v,and email is %v", tharun.Name, tharun.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
