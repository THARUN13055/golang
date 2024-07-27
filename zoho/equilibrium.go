// input: {1,3,4,2,2}
// output: 3 which is position

package main

import "fmt"

func main() {
	value := []int{1, 3, 4, 2, 2}
	lcount := 0
	rcount := 0
	position := 0
	for i := 0; i < len(value); i++ {
		rcount += value[i]
	}
	for j := 0; j < len(value); j++ {
		rcount -= value[j]
		if rcount == lcount {
			position = j + 1
		}
		lcount = lcount + value[j]
	}
	fmt.Print(position)
}
