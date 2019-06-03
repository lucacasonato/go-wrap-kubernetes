package main

import (
	"net/http"
)

func list(w http.ResponseWriter, r *http.Request) {
	// gets executed when all books are requested
}

func author(w http.ResponseWriter, r *http.Request) {
	// gets executed when all books from a specific author are requested
}

func main() {
	http.HandleFunc("/", list)
	http.HandleFunc("/author/{author}", author)
	http.ListenAndServe(":5000", nil)
}
