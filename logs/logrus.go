package main

import (
	"log"
	"os"
	"fmt"
)

func main(){
	file , err := os.OpenFile("./logfile.log",os.O_CREATE | os.O_APPEND | os.O_WRONLY , 0600 )
	if err != nil {
		fmt.Println(err)
	}
	
	log.SetOutput(file)

	log.Println("error is completely fine so no issues")
}