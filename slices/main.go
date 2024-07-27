package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruittype = []string{"apple", "tomoto", "banana"}
	fmt.Printf("what is the type of its %T\n", fruittype)
	fruittype = append(fruittype, "mango", "orange")
	fmt.Println(fruittype)

	fruittype = append(fruittype[1:3])
	fmt.Println(fruittype)

	highscores := make([]int, 4)
	highscores[0] = 2233
	highscores[1] = 434
	highscores[2] = 545
	highscores[3] = 636

	fmt.Println(highscores)

	highscores = append(highscores, 22342, 3423, 432, 42, 342, 42)

	fmt.Println(highscores)

	fmt.Println(sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println(highscores)

}
