package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	symbol := []string{"@", "$", "^"}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter the value: ")

	input, _ := reader.ReadString('\n')

	whitespace := strings.ReplaceAll(input, " ", "")

	s1 := strings.Split(whitespace, "")

	count := 0

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(symbol); j++ {
			if s1[i] == symbol[j] {
				count++
			}
		}
	}
	fmt.Println(count)
}
