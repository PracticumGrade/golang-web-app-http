package versionprinter

import (
	"fmt"
	"net/http"
)

type versionPrint struct {
	version   string
	buildTime string
}

func NewVersionPrinter(version, buildTime string) http.Handler {
	return &versionPrint{
		version:   version,
		buildTime: buildTime,
	}
}

func (vp *versionPrint) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "%s_%s\n", vp.version, vp.buildTime)
}
