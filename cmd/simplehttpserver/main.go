// авторское решение точки входа компилятора для всех модулей курса

package main

import (
	"fmt"
	"net/http"

	"github.com/PracticumGrade/web-app-http/internal/module1/http/printermux"
)

var (
	Version   = "vX.X.X"
	BuildTime = "YYYYMMDDhhmmss"
)

func main() {
	fmt.Printf("Version - %s, build time - %s.\n", Version, BuildTime)

	serveMux := http.NewServeMux()

	serveMux.Handle("/", printermux.New(Version, BuildTime))

	err := http.ListenAndServe(`:8080`, serveMux)
	if err != nil {
		panic(err)
	}
}
