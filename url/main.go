package main

import (
	"fmt"
	"net/http"
)

func main() {

	readwritefunc := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is a first command")
	}

	http.HandleFunc("/", readwritefunc)

	readwritefuncsecond := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is a second command")
	}

	http.HandleFunc("/cat", readwritefuncsecond)

	http.ListenAndServe(":5000", nil)
}
