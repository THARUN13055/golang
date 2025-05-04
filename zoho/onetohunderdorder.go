package main

import (
	"fmt"
	"sync"
)
var (
	mu sync.Mutex
	start = 1
	end = 100
	cond = sync.NewCond(&mu)
)

func printInOrder(i int, wg *sync.WaitGroup){
	defer wg.Done()

	mu.Lock()

	for i != start {
		cond.Wait()
	}

	fmt.Println(i)

	start++

	cond.Broadcast()

	mu.Unlock()
}


func main(){
	var wg sync.WaitGroup

	wg.Add(end)

	for i := 1; i <= end; i++ {
		go printInOrder(i, &wg)
	}
	wg.Wait()
	fmt.Println("All numbers printed in order from 1 to 100")
}

