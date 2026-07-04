package main

import "fmt"

func twoSum(num []int, target int) []int {
	for k := 0; k < len(num); k++ {
		for v := k + 1; v < len(num); v++ {
			if num[k]+num[v] == target {
				return []int{k, v}
			}
		}
	}
	return []int{-1, -1}
}

// func twoSum(num []int, target int) []int {
// 	maping := make(map[int]int)
// 	for i := 0; i < len(num); i++ {
// 		diff := target - num[i]
// 		if j, ok := maping[diff]; ok {
// 			return []int{j, i}
// 		}
// 		maping[num[i]] = i
// 	}
// 	return []int{-1, -1}
// }

func main() {
	value := twoSum([]int{1, 2, 4, 6, 3, 2}, 9)
	fmt.Println(value)
}
