package main

import "fmt"

type FrontUser struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type BackUser struct {
	Connect  string
	Tools    string
	Laungage string
	Codeline int
}

func main() {
	fmt.Println("struct in golang")

	// go has no inheritance and super or parent

	tharun := FrontUser{"tharun", "tharun@we323.go", true, 32}
	//fmt.Println(tharun)
	fmt.Printf("name is %v,and email is %v", tharun.Name, tharun.Email)

	backend := BackUser{"database", "kubernetes", "golang", 4000}
	//fmt.Println(backend)
	fmt.Printf("connection: %v, tools: %v, Laungage: %v, codeline: %v", backend.Connect, backend.Tools, backend.Laungage, backend.Codeline)
	fmt.Printf("Name: %v, Email: %v, Status: %v, Age: %v", tharun.Name, tharun.Email, tharun.Status, tharun.Age)
}
