package main

import (
	"fmt"
	
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var a int

	fmt.Println("enter the array size")
	fmt.Scan(&a)

	arrsize := make([]int, a)
	
	for i := 0; i < a; i++ {
		fmt.Scan(&arrsize[i])
	}

	for _, num := range arrsize {
		if isPrime(num) {
			fmt.Print(num)
		}
	}
	fmt.Println()

}
