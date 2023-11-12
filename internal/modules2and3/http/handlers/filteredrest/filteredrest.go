package filteredrest

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/PracticumGrade/web-app-http/internal/modules2and3/storage"
)

type (
	StorageI = storage.StorageI
	User     = storage.UserData

	usersManagement struct {
		usersStorage StorageI
	}
)

func NewFilteredRESTHandler(userStorage StorageI) http.Handler {
	return &usersManagement{usersStorage: userStorage}
}

func (um *usersManagement) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		um.getUsers(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func filterUsersByName(users []User, filter string) (filtered []User) {
	for _, usr := range users {
		if strings.Contains(usr.UserName, filter) {
			filtered = append(filtered, usr)
		}
	}

	return
}

func filterUsersByEmail(users []User, filter string) (filtered []User) {
	for _, usr := range users {
		if strings.Contains(usr.UserEmail, filter) {
			filtered = append(filtered, usr)
		}
	}

	return
}

func paginateUsers(users []User, page, limit int) (paginated []User) {
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	// Если начальный индекс выходит за границы, то возвращаем пустой список.
	if startIndex >= len(users) {
		paginated = []User{}

		return
	}

	// Если конечный индекс превышает общее количество записей, то устанавливаем его в конец.
	if endIndex > len(users) {
		endIndex = len(users)
	}

	paginated = append(paginated, users[startIndex:endIndex]...)

	return
}

func (um *usersManagement) getUsers(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	users, err := um.usersStorage.RecoverAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// Извлекаем параметры запроса.
	pageParam := r.FormValue("page")
	limitParam := r.FormValue("limit")
	nameParam := r.FormValue("name")
	emailParam := r.FormValue("email")

	if nameParam != "" {
		// Инициируем фильтр книг на основе 'name'.
		users = filterUsersByName(users, nameParam)
	}

	if emailParam != "" {
		// Инициируем фильтр книг на основе 'email'.
		users = filterUsersByEmail(users, emailParam)
	}

	if pageParam != "" && limitParam != "" {
		// Преобразуем параметры в целые числа и проводим проверку на допустимые значения параметров.
		page, err := strconv.Atoi(pageParam)
		if err != nil || page <= 0 {
			http.Error(w, "Invalid page parameter, must be positive integer", http.StatusBadRequest)

			return
		}

		limit, err := strconv.Atoi(limitParam)
		if err != nil || limit <= 0 {
			http.Error(w, "Invalid limit parameter, must be positive integer", http.StatusBadRequest)

			return
		}

		users = paginateUsers(users, page, limit)
	}

	if len(users) == 0 {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
