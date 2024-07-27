package main

import(
	"os"
	"fmt"
)

func main(){
	os.Setenv("name","tharun")
	fmt.Printf("%s envword",os.Getenv("name"))
}



