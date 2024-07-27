package main

import "fmt"

func main() {
	odd := [4]string{"tharun", "akash", "oman", "function"}
	for _, value := range odd {
		fmt.Printf("this is the value %s\n", value)
	}
}
