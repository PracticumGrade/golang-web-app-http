//file:<project>/cmd/<project>/main.go

package main

import (
	"net/http"

	"params-with-tests/internal/bookstore"
)

func main() {
	http.HandleFunc("/books", bookstore.BooksHandler)

	http.ListenAndServe(":8080", nil)
}
