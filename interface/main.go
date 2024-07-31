package main

import (
	"fmt"
	"os"
	"reflect"
)

type Mylogs interface {
	Value(string)
}

type Abcd struct{}

func (abcd Abcd) Value(m string) {
	fmt.Printf("this is tharun and its working %s", m)
}

type File struct {
	Filelogs *os.File
}

func (f *Filelogs) Value(fi string) {
	fmt.Fprint(f.File, fi)
}

func FileLoger(filename string) *Fileloges {
	files, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
}

func main() {
	var c Mylogs
	c = &Abcd{}
	c.Value("my name is tharun")
	fmt.Println(reflect.TypeOf(c))
	FileLoger()
}
