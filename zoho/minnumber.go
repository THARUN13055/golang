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
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	i1 := scanner.Text()
	sp := strings.Fields(i1)
	m := make([]int, len(sp))

	for i := 0; i < len(sp); i++ {
		sci, _ := strconv.Atoi(sp[i])
		m[i] = sci
	}
	for a:=0;a<len(m);a++{
		for b:=1;b<len(m);b++{
			if m[a:0] 
		}
	}
	
}
