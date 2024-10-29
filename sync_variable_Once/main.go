package main

import (
	"log"
	"time"
)

var value string

func data() {
	time.Sleep(100 * time.Millisecond)
	log.Println("working")
	value = "tharun"
}

func getdata(a string) string {
	if a != "" {
		data()
	}
	return value
}

func dosomthing(done chan struct{}) {
	getdata("")
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	for i := 0; i < 5; i++ {
		go dosomthing(done)
	}
	for i := 0; i < 5; i++ {
		<-done
	}
}
