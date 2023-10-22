//file:testwebapp/cmd/testwebapp/main.go

package main

import (
	"fmt"
	"net/http"

	"testwebapp/internal/handlers/simpleprinter"
)

var (
	Version   = "vX.X.X"
	BuildTime = "YYYYMMDDhhmmss"
)

func main() {
	fmt.Printf("Version - %s, build time - %s.\n", Version, BuildTime)

	mux := simpleprinter.NewSimplePrinterServeMux()

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
