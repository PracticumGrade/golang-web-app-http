// авторское решение из урока 2 темы 2.

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// RevolverError returns error with probability of 1/6%.
func RevolverError() error {
	rand.Seed(time.Now().UnixNano())
	probability := rand.Intn(6) + 1

	if probability == 6 {
		return errors.New("BOOM")
	}

	return nil
}

func checkIfBulletproofVestOn(headerValue string) bool {
	return headerValue == "on"
}

func RevolverErrorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed)

		return
	}

	response := "The revolver didn't shoot"

	err := RevolverError()
	if err != nil {
		if !checkIfBulletproofVestOn(r.Header.Get("X-Bulletproof-Vest")) {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		response = "Your bulletproof vest saved you"
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, response)
}

func main() {
	http.HandleFunc("/", RevolverErrorHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
