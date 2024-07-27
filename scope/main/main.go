package main

import (
	"myapp/packageone"
)

var myVar = "this is package level variable"

func main() {

	var blockVar = "this is block level variable"

	packageone.PrintMe(blockVar, myVar)

}
