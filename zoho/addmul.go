// input: 5   -> 120 34 22
// output: 3 12 4
// explination first add i[1] the number and multiple. compare add is > or mul is >.if add is > that is the answer.

package main

import "fmt"

func mul(input int) int {
	digit := []int{}

	for input > 0 {
		sepration := input % 10
		digit = append([]int{sepration}, digit...)
		input /= 10
	}
	val := 1
	for _, mul := range digit {
		val *= mul
	}
	return val
}

func add(input int) int {
	digit := []int{}

	for input > 0 {
		sepration := input % 10
		digit = append([]int{sepration}, digit...)
		input /= 10
	}
	val := 0
	for _, add := range digit {
		val += add
	}
	return val
}

func main() {
	var n int
	fmt.Print("Enter the array size:  ")
	fmt.Scanf("%d", &n)
	input := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &input[i])
	}
	answer := c
	for j := 0; j < n; j++ {
		sentmul := mul(input[j])
		sentadd := add(input[j])

		if sentadd > sentmul {
			answer = append(answer, sentadd)
		} else if sentadd < sentmul {
			answer = append(answer, sentmul)
		} else {
			answer = append(answer, sentadd)
		}

	}
	fmt.Println(answer)

}
