package main

import "fmt"

func main() {
	var arraysize int
	fmt.Println("Enter the Number or array")
	fmt.Scan(&arraysize)
	createarrsize := make([]int, arraysize)
	for i := 0; i < arraysize; i++ {
		fmt.Scan(&createarrsize[i])
	}
	var c []int
	var d []int
	for _, a := range createarrsize {
		if a != 0 {
			c = append(c, a)
		} else {
			d = append(d, a)
		}
	}
	var e []int
	e = append(c, d...)
	fmt.Println(e)
}
