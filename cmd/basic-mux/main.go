// авторское решение из урока 1 темы 3.

package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the About Page.")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", homeHandler)
	mux.HandleFunc("/about", aboutHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
