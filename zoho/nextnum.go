// Question 8
// Write a program that takes an integer M and M integer array elements as input. The program needs to find the minimum difference between two elements in the integer array. The program then needs to print all those pairs of elements that have the minimum difference. If more than one pair has the minimum difference, then the program should print the output in the ascending order, if an element exists in two or more pairs, then it should be printed two times or more.

// Example 1

// Input: 4
//            55 44 33 22
// Output:  22 33 33 44 44 55
// Explanation: The minimum difference between two elements is 11. Hence the pairs are printed in the ascending order. Here 33 and 44 appear in two different pairs; hence both are printed twice.
// Example 2

// Input: 5
//            1 99 22 44 1001
// Output:  1 22

package main

import (
	"bufio"
	"os"
	"strings"
)

func nextnumber(str string) []int {

	for _, r1 := range strings.Fields(str) {

	}

}

func scanner() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return nextnumber(scanner.Text())
}

func main() {
	a := scanner()

}
