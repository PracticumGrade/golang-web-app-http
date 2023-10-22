package printermux

import (
	"net/http"

	"github.com/PracticumGrade/web-app-http/internal/module1/http/handlers/randomnumberprinter"
	"github.com/PracticumGrade/web-app-http/internal/module1/randomprovider"

	"github.com/PracticumGrade/web-app-http/internal/module1/http/handlers/simpleprinter"
	"github.com/PracticumGrade/web-app-http/internal/module1/http/handlers/versionprinter"
)

func New(version, buildTime string) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", simpleprinter.MainPage)
	mux.HandleFunc("/api", simpleprinter.ApiPage)
	mux.Handle("/version", versionprinter.NewVersionPrinter(version, buildTime))
	mux.Handle("/randomnumber", randomnumberprinter.NewRandomNumberPrinter(randomprovider.NewRandomNumberProvider()))

	return mux
}
