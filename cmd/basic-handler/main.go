// авторское решение из урока 2 темы 1.

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	}))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
