// input this is will as

// input : i o a o s i c e o m n i u b t h

// output: is as this

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter the first line")
	scanner.Scan()
	i1 := scanner.Text()
	fmt.Println("enter the second line")
	scanner.Scan()
	i2 := scanner.Text()
	fmt.Println("here now in execution process")
	var myarray string

	inp1 := strings.Split(i1, " ")
	inp2 := strings.Split(i2, "")

	i := 0
	for i < len(inp1) {
		spliti1 := strings.Split(inp1[i], "")
		a := 0
		match := 0
		for a < len(spliti1) {
			b := 0
			for b < len(inp2) {
				if spliti1[a] == inp2[b] {
					match += 1
					break
				}
				b++
			}
			a++
		}
		if len(inp1[i]) == match {
			fmt.Println(inp1[i])
		}
		i++

	}
	fmt.Println(myarray)

}
