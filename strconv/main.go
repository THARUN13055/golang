package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main()  {
	a := "12";
	intVar,err := strconv.Atoi(a)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
	fmt.Printf("a value is : %v and type is %T \n", intVar, intVar)

	pas,err := strconv.ParseInt(a,0,16)
	fmt.Println(pas,err,reflect.TypeOf(pas))

}