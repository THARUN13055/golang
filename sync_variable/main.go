package main

import (
	"fmt"
	"sync"
)

func main() {
	var value int
	value = 100000000
	temp := 0
	done := make(chan struct{})
	var mux sync.Mutex
	go func() {
		for i := 0; i < value; i++ {
			mux.Lock()
			temp++
			mux.Unlock()
		}
		done <- struct{}{}
	}()
	go func() {
		for i := 0; i < value; i++ {
			mux.Lock()
			temp--
			mux.Unlock()
		}

		done <- struct{}{}
	}()
	<-done
	<-done
	fmt.Println(temp)
}
