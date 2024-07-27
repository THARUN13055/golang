// package main

// func main() {

// 	value := make(chan int, 1)

// 	value <- 90

// 	n := <-value

// 	fmt.Printf("%d", n)
// }

// sent and recive channal

// type abcd struct {
// 	name string
// 	id   int
// }

// func sentandrecive(ch chan abcd, a abcd) {
// 	ch <- a
// }
// func main() {
// 	a := abcd{"tharun", 23}
// 	cha := make(chan abcd)
// 	go sentandrecive(cha, a)
// 	name := (<-cha).name
// 	fmt.Println(name)

// }

//-----------------

// func main() {
// 	a := time.Now()
// 	done := make(chan int)

// 	go func() {
// 		for i := 0; i < 10000; i++ {
// 			done <- i
// 		}
// 		close(done)

// 	}()
// 	for n := range done {
// 		fmt.Println(n)
// 	}
// 	// for i := 0; i < 1000; i++ {
// 	// 	fmt.Println(i)
// 	// }
// 	c := time.Since(a)
// 	fmt.Println(c)

// }
 