package main

import (
	"encoding/json"
	"os"
)

type data struct {
	Name  string `json: "name"`
	Email string `json: "email"`
}


func main(){
  file,err := os.Open("test.json")
  if err != nil {
    fmt.Println("files is not opening")
  }

  defer file.Close()

  decode := json.Decoder(file)
  

}