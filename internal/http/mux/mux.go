package mux

import (
	"net/http"

	"github.com/PracticumGrade/web-app-http/internal/modules2and3/http/handlers/filteredrest"

	"github.com/PracticumGrade/web-app-http/internal/modules2and3/storage"

	"github.com/PracticumGrade/web-app-http/internal/modules2and3/http/handlers/rest"

	"github.com/PracticumGrade/web-app-http/internal/module1/http/handlers/randomnumberprinter"
	"github.com/PracticumGrade/web-app-http/internal/module1/randomprovider"

	"github.com/PracticumGrade/web-app-http/internal/module1/http/handlers/simpleprinter"
	"github.com/PracticumGrade/web-app-http/internal/module1/http/handlers/versionprinter"
)

func New(version, buildTime string) *http.ServeMux {
	fileStorage := storage.New(".")
	m := http.NewServeMux()

	m.HandleFunc("/printer/hello", simpleprinter.MainPage)
	m.HandleFunc("/printer/api", simpleprinter.ApiPage)
	m.Handle("/printer/version", versionprinter.NewVersionPrinter(version, buildTime))
	m.Handle("/printer/randomnumber",
		randomnumberprinter.NewRandomNumberPrinter(randomprovider.NewRandomNumberProvider()))

	m.Handle("/restapi/users", rest.NewRESTHandler(fileStorage))
	m.Handle("/restapi/users/filtered", filteredrest.NewFilteredRESTHandler(fileStorage))

	return m
}
