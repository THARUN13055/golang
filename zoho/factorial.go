package main

import (
	"fmt"
)

func zero(mul int) int {
	count := 0
	flag := 0
	for mul != 0 {
		flag = mul / 5
		count += flag
		mul = mul / 5

	}
	return count
}
func main() {
	var a int
	fmt.Scanf("%d", &a)

	result := zero(a)
	fmt.Println(result)
}
