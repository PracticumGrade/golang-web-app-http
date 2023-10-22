//file:testwebapp/internal/handlers/simpleprinter/mux.go

package simpleprinter

import "net/http"

func NewSimplePrinterServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainPage)
	mux.HandleFunc("/api", apiPage)

	return mux
}
