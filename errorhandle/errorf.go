package main

import (
	"fmt"
)

func checkvalue(n int) error {
	if n%2 == 0 {
		return fmt.Errorf("Only odd number is allow and the value is: %d", n)
	}
	return nil
}

func checkerror(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("the code is executed success")
}

func main() {
	err := checkvalue(2)
	checkerror(err)
}
