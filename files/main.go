package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "this will we need to push in the file- learn cloud"

	file, err := os.Create("./myfile.txt")

	io.WriteString(file, content)
	checkerror(err)

	length, err := io.WriteString(file, content)
	checkerror(err)

	fmt.Println("length is ", length)
	defer file.Close()
	readfile("./myfile.txt")

}

func readfile(filename string) {

	databyte, err := ioutil.ReadFile(filename)

	checkerror(err)
	fmt.Println("text data inside the file is \n", databyte)
	fmt.Println(string(databyte))

}

func checkerror(err error) {
	if err != nil {
		panic(err)
	}
}
