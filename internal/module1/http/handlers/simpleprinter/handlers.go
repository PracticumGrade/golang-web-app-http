package simpleprinter

import (
	"fmt"
	"net/http"
)

func MainPage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func ApiPage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "This is \"/api\" page.")
}
