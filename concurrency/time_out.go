package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan int)

	go sent(ch)

	select {
	case mgs := <- ch:
		fmt.Println(mgs)
	case <- time.After(1 * time.Second):
		fmt.Println("session timeout as a given second")
	}
}

func sent(ch1 chan int){
	// 	
}