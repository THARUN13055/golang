/*
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-->")

		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(userInput)
		}

	}
}
*/

package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {

		_ = keyboard.Close()

	}()

	fmt.Println("press any key of the keyboard. Press esc to quit.")
	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if key != 0 {
			fmt.Println("you pressed", char, key)
		} else {
			fmt.Println("you pressed", char)
		}
		if key == keyboard.KeyEsc {
			break
		}
	}
}
