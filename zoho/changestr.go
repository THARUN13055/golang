package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "tharun is a good boy"
	newstring := "**"
	replacevalue := "o"
	// replace := strings.ReplaceAll(a, "o", newstring)
	// fmt.Println(replace)

	split := strings.Fields(a)
	count := ""
	for i, replace := range split {
		if i > 0 {
			count += " "
		}
		if replace != replacevalue {
			count += replace
		} else {
			count += newstring
		}
	}

	fmt.Println(count)
}
