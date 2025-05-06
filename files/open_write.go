package main

import (
	"fmt"
	"os"
)

func main() {
	// open the file here if the file is not there it will create os.O_Create

	openfile, err := os.OpenFile("./tharun.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer openfile.Close()

	// Write the string

	values := "tharun is good boy"

	_, err = openfile.WriteString(values)
	if err != nil {
		fmt.Println(err)
		return
	}

}
