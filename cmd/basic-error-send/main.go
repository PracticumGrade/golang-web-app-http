// авторское решение из урока 1 темы 2.

package main

import (
	"net/http"
)

func badResponse(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("X-Custom-Header", "custom-header-value")
	http.Error(w, "error description", http.StatusInternalServerError)
}

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(badResponse))
	if err != nil {
		panic(err)
	}
}
