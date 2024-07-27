// tharun
// thrn

// thaaruun
//thaaruun

package main

import (
	"fmt"
	"strings"
)

func main() {
	var value string
	fmt.Scanln(&value)
	vo := []string{"a", "e", "i", "o", "u"}
	var out []string
	i1 := strings.Split(value, "")
	for i := 0; i < len(i1); i++ {
		match := false
		for j := 0; j < len(vo); j++ {
			if i1[i] == vo[j] {
				match = true
				if i1[i] == i1[i+1] {
					out = append(out, i1[i], i1[i+1])
					i++
				}
			}

		}
		if !match {
			out = append(out, i1[i])
		}
	}
	fmt.Print(out)

}
