// package main

// import "fmt"

// import "fmt"

// func main() {
// 	a := 4
// 	normal(a)
// 	pointerl1(&a)
// 	fmt.Printf("this is the main a: %v", &a)
// }

// func normal(v int) {
// 	fmt.Println(&v)
// }

// func pointerl1(v *int) {
// 	fmt.Println("address because v have pointed to address: ", v, "here * mension the value which inside in addr: ", *v)
// }

// type 2

// type val struct {
// 	m int
// }

// type val2 struct {
// 	c int
// }

// func pointer() *val {
// 	a := val{m: 3}
// 	fmt.Printf("%p\n", &a)
// 	return &a
// }

// func poin() *int {
// 	m := 40
// 	fmt.Printf("%p\n", &m)
// 	return &m
// }

// func main() {
// 	fmt.Println(pointer())
// 	p := poin()
// 	q := &p

// 	fmt.Println(p)
// 	fmt.Printf("%T", &p)
// 	fmt.Println(*q)

// }

// package main

// import "fmt"

// func main() {
// 	a := 5
// 	c := &a
// 	b := *c
// 	fmt.Println(b)
// }

package main

import "fmt"

func main() {
	var a int = 4
	fmt.Println(&a)
	var prt *int = &a
	fmt.Println(prt)

	fmt.Println(*prt)

}
