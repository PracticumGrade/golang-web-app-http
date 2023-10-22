package randomnumberprinter

import (
	"fmt"
	"net/http"

	"github.com/PracticumGrade/web-app-http/internal/module1/randomprovider"
)

type (
	RandomNumberProvider = randomprovider.RandomNumberProvider

	randNumbPrint struct {
		r RandomNumberProvider
	}
)

func NewRandomNumberPrinter(provider RandomNumberProvider) http.Handler {
	return &randNumbPrint{r: provider}
}

func (rnp *randNumbPrint) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "%d\n", rnp.r.GetRandomInt())
}
