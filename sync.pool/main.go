package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("create the new slice")
			return make([]byte, 1024)
		},
	}
	// Get a slice from the pool
	obj := pool.Get().([]byte)
	fmt.Printf("Got a slice of length %d\n", len(obj))

	// pool.Put(obj)

	reusedObj := pool.Get().([]byte)
	fmt.Printf("Got a reused slice of length %d\n", len(reusedObj))

	pool.Put(reusedObj)
}
