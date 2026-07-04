package main

import (
	"fmt"
	"strings"
	"strconv"
)

func arr(val []int) int {
	result := 0

	// take each number first
	for i := 0; i<len(val); i++ {
	
		// get each number
		numb := val[i]
		// convert number as string
		convertstring := strconv.Itoa(numb)
		// Now we need to split each string
		digits := strings.Split(convertstring,"")

		// now we need to changes it as int and added it
		sum := 0
		for _,values := range digits{
			digt, _ := strconv.Atoi(values)
			sum += digt
		}
		result += sum
		fmt.Printf("getting %d\n",sum)
	}
	return result
}

func main (){
	a := []int{12,33,21}
	fmt.Printf("the value is : %d ", arr(a))
}
