//Input [0,1,1,2]
//output [0,1,2,_]

package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 1, 2, 3, 3, 5, 5, 5, 6, 1, 4} // Input array
	uniqueMap := make(map[int]bool)                // Map to track unique elements
	uniqueArray := []interface{}{}                 // Array to store unique elements with "_" as placeholders

	// Loop through the input array
	for _, value := range a {
		// If the value hasn't been added yet, add it to uniqueArray
		if !uniqueMap[value] {
			uniqueArray = append(uniqueArray, value)
			uniqueMap[value] = true // Mark this value as seen
		}
	}

	// Fill the array with "_" to reach the original length
	for len(uniqueArray) < len(a) {
		uniqueArray = append(uniqueArray, "_")
	}

	// Output the result
	fmt.Println(uniqueArray) // Expected output: [0, 1, 2, _]
}
