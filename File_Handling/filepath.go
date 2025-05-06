package main

import (
  "fmt"
  "os"
)

func main (){
  f,err := os.Stat("/home/work/abcd.sh")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(f.Name())
  fmt.Println(f.Size())
}
