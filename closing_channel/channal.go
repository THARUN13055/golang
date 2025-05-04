package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup

	wg.Add(2)

	ch := make(chan int)

	go sell(ch, &wg)
	go buy(ch, &wg)
	wg.Wait()

}

func sell(ch chan int , wg *sync.WaitGroup) {
	ch <- 10
	ch <- 11

	// printing the values
	fmt.Println("Values has been sell")
	close(ch)
	wg.Done()
}


func buy(ch chan int, wg *sync.WaitGroup){
	fmt.Println("waiting for the chan get data")

	for val := range ch {
		fmt.Println("the value is: ",val)
	}
	wg.Done()
}