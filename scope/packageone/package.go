package packageone

import "fmt"

var Packagevar = "this is package level variable12324"

func PrintMe(s1, s2 string) {

	fmt.Println(s1, s2, Packagevar)

}
