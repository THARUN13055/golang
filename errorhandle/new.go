package main

import (
	"errors"
	"fmt"
)

func checkvalue(n int) error {
	if n%2 == 0 {
		return errors.New("Only odd number is allow")
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
