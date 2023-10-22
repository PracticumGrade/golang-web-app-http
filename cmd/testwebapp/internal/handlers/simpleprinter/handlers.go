//file:testwebapp/internal/handlers/simpleprinter/handlers.go

package simpleprinter

import (
	"fmt"
	"net/http"
)

func mainPage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func apiPage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "This is \"/api\" page.")
}
