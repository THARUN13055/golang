/*

//helloworld

package main

import "fmt"

func main() {
	helloname("this is hello world")
}

func helloname(say string) {
	fmt.Println(say)
}

//variable

package main

import "fmt"

func main() {
	a := 4
	fmt.Printf("%d", a)
}

//ifelse

package main

import "fmt"

func main() {

	var a int

	fmt.Println("enter the number: ")
	fmt.Scanf("%d", &a)
	if a%2 == 0 {
		fmt.Printf("%d is a even number", a)

	} else {
		fmt.Printf("%d is the odd number ", a)
	}
}

//ifelse

package main

import "fmt"

func main() {

	var a int

	fmt.Printf("enter the number: ")
	fmt.Scanf("%d", &a)

	n := a / 2

	if a%2 == 0 {
		fmt.Printf("%d is even , %d time it will divide", a, n)

	} else {
		fmt.Printf("%d is odd,%d time it will divide", a, n)
	}
}

// if else fizzbizz

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Printf("enter the number: ")
	fmt.Scanf("%d", &num)
	fmt.Printf("%s", fizzbizz(num))

}

func fizzbizz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "fizzbizz"
	} else if n%3 == 0 {
		return "fizz"
	} else if n%5 == 0 {
		return "bizz"
	}
	return strconv.Itoa(n)
}


package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Printf("enter the number: ")
	fmt.Scanf("%d", &n)
	fmt.Printf("%s", secname(n))

}

func secname(s int) string {
	if s%2 == 0 {
		return "its correct"
	} else {
		return "its false"
	}

}


package main

import "fmt"

func main() {
	var a int
	fmt.Printf("enter the number of watercane: ")
	fmt.Scanf("%d", &a)
	fmt.Printf("%s", watercane(a))

}
func watercane(a int) string {
	if a > 50 {
		return " this is over capacity"
	} else {
		return "this is medium capacity"
	}
}


// Error messaging while no output given

package helloweb

import (
	"errors"
	"fmt"
)

func hello(name string) (string, error) {
	if name == "tharun" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("hi %v. welcome", name)
	return message, nil
}


package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "press enter to ready"

func main() {
	fno := 2
	sno := 5
	tno := 7
	var ans int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("guess the number 1 to 10", fno, prompt)
	reader.ReadString('\n')
	fmt.Println("subract the number", sno, prompt)
	reader.ReadString('\n')
	fmt.Println("divid the number", tno, prompt)
	reader.ReadString('\n')

	ans = fno*sno - tno

	fmt.Println("the answer is ", ans)

}


package main

import (
	"fmt"
)

func main() {
	var a byte
	fmt.Println("guess the name:(y/n) ")
	fmt.Scanf("%c", &a)
	// go doc unicode
	// now we need to convert the rune (unicode character) but we are using byte which means we need to change as ASCII value
	//a = byte(unicode.ToLower(rune(a)))
	//a = byte(unicode.ToUpper(rune(a)))
	switch a {
	case 'y':
		fmt.Println("thanks")
	case 'n':
		fmt.Println("no thanks")
	default:
		fmt.Println("bye")
	}
}




package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}



// here we are declearing the single function value
package main

import "fmt"

func main() {
	var a = "this is string"

	fmt.Println(a)
	myFunc()
}
func myFunc() {
	var a = "this is also string!"
	fmt.Println(a)
}



package main

import "fmt"

// we can be used it for multi time

var a = "this is apply for all"

func myFunc() {
	fmt.Println(a)
}

func main() {
	fmt.Println(a)
	myFunc()
}




package main

import (
	"paclageone"
)

var myVar = "this is package level variable"

func main() {

	var blockVar = "this is block level variable"

	paclageone.PrintMe(blockVar, myVar)

}

//Pointer

package main

import "fmt"

func main() {

	a := 30

	b := "string"

	var Pointer = &a

	var po = &b

	fmt.Println("int value:", *Pointer)
	fmt.Println("int  address:", Pointer)
	fmt.Println("--------------------")
	fmt.Println("string value", *po)
	fmt.Println("string address", po)
}


//swapping of two strings
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("tharun", "kumar")
	fmt.Println(a, b)
}
*/