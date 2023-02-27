package main

import "fmt"

func main() {
	fmt.Println("this is map")

	language := make(map[string]string)

	language["JS"] = "javascript" //here js is a key and javascript is a value
	language["py"] = "python"
	language["RB"] = "ruby"

	fmt.Println("list of all laungage", language)

	fmt.Println("only show the fullform for py", language["py"])

	delete(language, "RB")
	fmt.Println("list of all laungage", language)
	//loops

	for key, value := range language {
		fmt.Printf("the key is %v, and the value is %v\n", key, value)
	}

}
