package main

import "fmt"

func comming(in chan int) {
	for i := 0; i < 5; i++ {
		in <- i
	}
	close(in)
}

func going(out chan int, done chan struct{}) {
	for i := range out {
		fmt.Println(i)
	}
	done <- struct{}{}
}

func storing(in chan int, out chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	done := make(chan struct{})
	go comming(chan1)
	go storing(chan1, chan2)
	go going(chan2, done)
	<-done
}
