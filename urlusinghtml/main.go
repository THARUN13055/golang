package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("content-type", "text/plain-text")

		w.Header().Set("content-length", "40")

		fmt.Fprintf(w, "<h1>this is a first command</h1>")
	})

	http.HandleFunc("/cat", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "cat.html")
	})
	http.ListenAndServe(":5050", nil)
}
