package main

import "fmt"

func main() {
	fmt.Println("this is ifelse statement")

	loginincount := 34
	var result string

	if loginincount < 20 {
		result = "tharun"

	} else if loginincount > 32 {
		result = "no"
	} else {
		result = "bad"
	}
	fmt.Println(result)

}
