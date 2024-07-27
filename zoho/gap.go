// input : 1 4 5 2 4 5 3
// output : 13 11

package main

import "fmt"

func main() {
	fmt.Print("Enter the value: ")
	var gap int
	var n int
	fmt.Scanf("%d", &n)
	makes := make([]int, n)
	for m := 0; m < n; m++ {
		fmt.Scan(&makes[m])
	}

	fmt.Print("Gap of the value: ")
	fmt.Scanf("%d", &gap)
	count := 0
	for i := 0; i < len(makes); i++ {
		count += makes[i]
		i = i + gap - 1
	}
	scount := 0
	for i := 1; i < len(makes); i++ {
		scount += makes[i]
		i = i + gap - 1
	}
	fmt.Println(count)
	fmt.Println(scount)
}
