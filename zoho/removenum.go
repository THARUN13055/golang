package main

import (
	"fmt"
	"strconv"
	"strings"
)

func answer(input1 string) string {
	field := strings.Fields(input1)

	store := make([]string, len(field))

	for _, ran1 := range field {
		//field
		s2 := strings.Split(ran1, "")
		var app strings.Builder
		conv := 0
		for _, ran2 := range s2 {

			if !(ran2 >= "0" && ran2 <= "9") {
				app.WriteString(ran2)
			} else {
				conv, _ = strconv.Atoi(ran2)
			}

		}
		store[conv] = app.String()
	}
	fmt.Print(store)
	return ""
}

func main() {
	input := "t2o j3oin 4WonderBiz I0 Technolog5ies wan1t"
	output := answer(input)
	fmt.Println(output)
}


