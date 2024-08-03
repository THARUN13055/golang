package main

import (
	"fmt"
)

// func abcd(returning chan bool) {
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("hi i am tharun")
// 	returning <- true
// }

// func main() {
// 	makingchannal := make(chan bool)
// 	go abcd(makingchannal)
// 	<- makingchannal
// }

//-----------------------------------
// func main() {
// 	message := make(chan int)

// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		message <- 1
// 	}()

// 	go func() {
// 		time.Sleep(8 * time.Second)
// 		message <- 2
// 	}()

// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		message <- 3
// 	}()

// 	go func() {
// 		fmt.Println(4 * time.Second)
// 		message <- 4
// 	}()

// 	go func() {
// 		time.Sleep(6 * time.Second)
// 		message <- 5
// 	}()

// 	// Receiver goroutine to handle incoming messages
// 	go func() {
// 		for i := range message {
// 			fmt.Println(i)
// 		}
// 	}()

// 	// Wait for all goroutines to finish
// 	time.Sleep(15 * time.Second) // Ensure the channel is processed

// 	// Close the channel
// 	close(message)

// 	// Give some time for the receiver to finish processing
// 	time.Sleep(2 * time.Second)
// }
//-------------------------------------------

func main() {
	creating := make(chan int)
	go sending(creating)
	for a := range creating {
		fmt.Println(a)
	}
}

func sending(out chan int) {
	for i := 0; i < 5; i++ {
		out <- i
	}
	defer close(out)
}
