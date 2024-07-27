package main

import (
	"fmt"
)

type Student struct {
	name string
	id   int
}

type Child struct {
	bond string
	year int
}

type Parent struct {
	age    int
	gender string
}
type Animals struct {
	name string
	age  int
}

func (s Student) returnval(p Parent, c Child, a Animals) {
	fmt.Printf("this is the value: %s %d\n", s.name, s.id)
	fmt.Printf("this is child value: %s %d\n", c.bond, c.year)
	fmt.Printf("this is parent value: %d %s\n", p.age, p.gender)
	fmt.Printf("this is animals value: %s %d\n", a.name, a.age)
}

func main() {
	book := Student{"thaurn", 12}
	chi := Child{"male", 2020}
	par := Parent{34, "Female"}
	ani := Animals{"lion", 45}
	book.returnval(par, chi, ani)
}
