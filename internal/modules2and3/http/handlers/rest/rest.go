package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/PracticumGrade/web-app-http/internal/modules2and3/storage"
)

type (
	StorageI = storage.StorageI
	User     = storage.UserData

	usersManagement struct {
		usersStorage StorageI
	}
)

func NewRESTHandler(userStorage StorageI) http.Handler {
	return &usersManagement{usersStorage: userStorage}
}

func (um *usersManagement) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		um.getUsers(w, r)
	case http.MethodPost:
		um.createUser(w, r)
	case http.MethodPut:
		um.updateUser(w, r)
	case http.MethodDelete:
		um.deleteUser(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (um *usersManagement) getUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := um.usersStorage.RecoverAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if len(users) == 0 {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (um *usersManagement) createUser(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	id, err := um.usersStorage.Store(newUser)
	if err != nil {
		if errors.Is(err, storage.ErrAlreadyExists) {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	newUser.UserID = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (um *usersManagement) updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	var userToUpdate User

	err := json.NewDecoder(r.Body).Decode(&userToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = um.usersStorage.Update(userToUpdate)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)

			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userToUpdate)
}

func (um *usersManagement) deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	var userToDelete User

	err := json.NewDecoder(r.Body).Decode(&userToDelete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = um.usersStorage.Delete(userToDelete)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)

			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}
