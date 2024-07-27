// Input: (9*(7-2)*(1*5)
// Output: 0

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	i1, _ := reader.ReadString('\n')

	s1 := strings.Split(i1, "")

	fmt.Println(s1)
	i := 0
	count := 0
	value := 0
	fmt.Println(len(s1))
	for i < len(s1) {
		if "(" == s1[i] {
			count += 1
		} else if ")" == s1[i] {
			value += 1
		}
		i++
	}
	a := count - value
	// main code over
	//------------------
	if a < 0 {
		fmt.Print("0")
	} else if a > 0 {
		fmt.Print("0")
	} else {
		fmt.Print("1")
	}
}
