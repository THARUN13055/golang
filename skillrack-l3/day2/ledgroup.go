// Question

// we need to make group of colors

// inputs:
// red,blue,green,blue

// output:
// 5

// explination:
// red,blue,green,blue,(blue,green,blue)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the value in single line")
	input, _ := reader.ReadString('\n')
	a := strings.Fields(strings.TrimSpace(input))
	// a := [7]string{"yello","red","yello","green","blue","yello","green"}
	sum := 0
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				sum++
			}
		}
	}
	final := len(a) + sum
	fmt.Printf("the group of value is: %v ", final)
}
