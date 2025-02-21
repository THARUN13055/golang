package main

import (
	"fmt"
	"io"
	"os"
)

func main(){

	// Here i am creating the file

	file , err := os.Create("test.txt")
	if err != nil {
		fmt.Println("the file is not creted",err)
		return
	}
	defer file.Close()

	// Now i need to write the data using io.Writter
	// Write a string to the file using io.Writer

	var writter io.Writer = file

	// Now we write the string into the file

	_, err = writter.Write([]byte("tharun"))
	if err != nil {
		fmt.Println("it cant able to write",err)
		return
	}

}