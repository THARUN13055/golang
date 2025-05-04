package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup

	wg.Add(2)

	go leaks(&wg)

	wg.Wait()
}

func leaks(s *sync.WaitGroup){
	ch := make(chan int)

	go func(){
		val := <- ch
		fmt.Println("the value is ", val)
		s.Done()
    	
	}()

	fmt.Println("exiting the leak method")
	s.Done()
}