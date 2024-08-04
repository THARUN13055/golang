package main

import (
	"fmt"
	"sync"
)
// actually lock and unlock is user for the changes which will not be occurs semiltaniously because here the example here i need to incress one and delete one if the go routine run parally it get error
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
