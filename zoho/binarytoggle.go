package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	dic := 10
	dectobi := strconv.FormatInt(int64(dic), 2)
	fmt.Println("this is binary number: %d", dectobi)
	split := strings.Split(dectobi, "")
	for i := 0; i < len(split)/2; i++ {
		j := len(split) - 1 - i
		split[i], split[j] = split[j], split[i]
	}
	reversedBinary := strings.Join(split, "")
	fmt.Println(reversedBinary)

}
