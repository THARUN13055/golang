package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	Name   string `json:"Name"`
	Friend string `json:"Friend"`
}


func studenthandle(w http.ResponseWriter, r *http.Request) {
	s := student{
		Name:   "tharun",
		Friend: "mass",
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func firstone(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is the display")
}
func secondone(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func main() {
	http.HandleFunc("/one", firstone)
	http.HandleFunc("/second", secondone)
	http.HandleFunc("/student", studenthandle)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("There port 8080 contain error", err)
	}
}
