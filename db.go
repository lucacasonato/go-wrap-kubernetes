package main

import (
	"time"

	"github.com/lucacasonato/wrap/filter"

	"github.com/lucacasonato/wrap"
)

// Book info
type Book struct {
	Name   string
	Author string
	ISBN   string
}

var books *wrap.Collection

func init() {
	// initalize database
	client, err := wrap.Connect("mongodb://mongo-0.mongo,mongo-1.mongo:27017", 5*time.Second) // create a new client with a 5 second timeout
	if err != nil {
		panic(err)
	}

	db := client.Database("main")
	books = db.Collection("books")
	err = books.Delete() // remove old data from collection
	if err != nil {
		panic(err)
	}

	// add some books
	_, err = books.Add(&Book{
		Name:   "The Go Programming Language",
		Author: "Alan A. A. Donovan",
		ISBN:   "0134190440",
	})
	if err != nil {
		panic(err)
	}

	_, err = books.Add(&Book{
		Name:   "Kubernetes: The Complete Guide To Master Kubernetes",
		Author: "Joseph D. Moore",
		ISBN:   "1096165775",
	})
	if err != nil {
		panic(err)
	}

	_, err = books.Add(&Book{
		Name:   "Mastering MongoDB 4.x: Expert techniques to run high-volume and fault-tolerant database solutions using MongoDB 4.x",
		Author: "Alex Giamas",
		ISBN:   "1789617871",
	})
	if err != nil {
		panic(err)
	}

	// create an index
	err = books.CreateIndex(map[string]wrap.Index{
		"author": wrap.TextIndex,
	})
	if err != nil {
		panic(err)
	}
}

func allBooks() ([]Book, error) {
	iterator, err := books.
		All().
		DocumentIterator()
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	b := []Book{}

	for iterator.Next() {
		book := Book{}

		err = iterator.DataTo(&book)
		if err != nil {
			return nil, err
		}

		b = append(b, book)
	}

	return b, nil
}

func authorBooks(author string) ([]Book, error) {
	iterator, err := books.
		Where(filter.TextSearch(author)).
		DocumentIterator()
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	b := []Book{}

	for iterator.Next() {
		book := Book{}

		err = iterator.DataTo(&book)
		if err != nil {
			return nil, err
		}

		b = append(b, book)
	}

	return b, nil
}
