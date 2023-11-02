// авторское решение точки входа компилятора для всех модулей курса

package main

import (
	"fmt"
	"net/http"

	"github.com/PracticumGrade/web-app-http/internal/http/mux"
)

var (
	Version   = "vX.X.X"
	BuildTime = "YYYYMMDDhhmmss"
)

func main() {
	fmt.Printf("Version - %s, build time - %s.\n", Version, BuildTime)

	err := http.ListenAndServe(`:8080`, mux.New(Version, BuildTime))
	if err != nil {
		panic(err)
	}
}
