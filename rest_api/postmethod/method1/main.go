package main

import (
	"encoding/json"
	"net/http"
)

type move struct {
	Title       string `json:"Title"`
	Author      string `json:"Author"`
	Description string `json:"Description"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	post := []move{
		{"mass", "tharun", "this is tharun mass story"},
		{"losu", "sanjay", "this is sanjay story"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func main() {
	http.HandleFunc("/move", PostHandler)
	http.ListenAndServe(":5050", nil)
}
