package main

import (
	"fmt"
	"net/http"
)

// Writting the middleware function

// type HandlerFunc func(http.ResponseWriter, *http.Request)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.FormValue("user")
		password := r.FormValue("password")

		if user != "user" || password != "password" {
			http.Error(w, "The user and password is invalid", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// A list of middleware functions
var middleware = []func(http.Handler) http.Handler{
	authMiddleware,
}

// Final handler
func work(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Work done")
}

func main() {

	var h http.Handler = http.HandlerFunc(work)
	for _, m := range middleware {
		h = m(h)
	}
	http.HandleFunc("/work", h.ServeHTTP)
	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
