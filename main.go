package main

import (
	"encoding/json"
	"net/http"
)

func list(w http.ResponseWriter, r *http.Request) {
	books, err := allBooks()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func author(w http.ResponseWriter, r *http.Request) {
	author := r.FormValue("author")

	books, err := authorBooks(author)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func main() {
	http.HandleFunc("/", list)
	http.HandleFunc("/author", author)
	http.ListenAndServe(":5000", nil)
}
