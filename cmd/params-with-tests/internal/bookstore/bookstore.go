//file:<project>/internal/bookstore/bookstore.go

package bookstore

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books []Book

func init() {
	// Наполняем базу данных 50 книгами.
	books = []Book{
		{1, "Book 1", "Author 1", 19.99},
		{2, "Book 2", "Author 2", 24.99},
		{3, "Book 3", "Author 3", 14.99},
		{4, "Book 4", "Author 1", 29.99},
		{5, "Book 5", "Author 2", 9.99},
		{6, "Book 6", "Author 3", 17.99},
		{7, "Book 7", "Author 1", 21.99},
		{8, "Book 8", "Author 2", 13.99},
		{9, "Book 9", "Author 3", 28.99},
		{10, "Book 10", "Author 1", 12.99},
		{11, "Book 11", "Author 2", 23.99},
		{12, "Book 12", "Author 3", 18.99},
		{13, "Book 13", "Author 1", 22.99},
		{14, "Book 14", "Author 2", 9.99},
		{15, "Book 15", "Author 3", 14.99},
		{16, "Book 16", "Author 1", 27.99},
		{17, "Book 17", "Author 2", 21.99},
		{18, "Book 18", "Author 3", 19.99},
		{19, "Book 19", "Author 1", 25.99},
		{20, "Book 20", "Author 2", 11.99},
		{21, "Book 21", "Author 3", 26.99},
		{22, "Book 22", "Author 1", 22.99},
		{23, "Book 23", "Author 2", 17.99},
		{24, "Book 24", "Author 3", 29.99},
		{25, "Book 25", "Author 1", 19.99},
		{26, "Book 26", "Author 2", 24.99},
		{27, "Book 27", "Author 3", 13.99},
		{28, "Book 28", "Author 1", 27.99},
		{29, "Book 29", "Author 2", 31.99},
		{30, "Book 30", "Author 3", 11.99},
		{31, "Book 31", "Author 1", 22.99},
		{32, "Book 32", "Author 2", 14.99},
		{33, "Book 33", "Author 3", 19.99},
		{34, "Book 34", "Author 1", 24.99},
		{35, "Book 35", "Author 2", 12.99},
		{36, "Book 36", "Author 3", 28.99},
		{37, "Book 37", "Author 1", 23.99},
		{38, "Book 38", "Author 2", 17.99},
		{39, "Book 39", "Author 3", 21.99},
		{40, "Book 40", "Author 1", 16.99},
		{41, "Book 41", "Author 2", 22.99},
		{42, "Book 42", "Author 3", 29.99},
		{43, "Book 43", "Author 1", 19.99},
		{44, "Book 44", "Author 2", 24.99},
		{45, "Book 45", "Author 3", 17.99},
		{46, "Book 46", "Author 1", 21.99},
		{47, "Book 47", "Author 2", 13.99},
		{48, "Book 48", "Author 3", 28.99},
		{49, "Book 49", "Author 1", 27.99},
		{50, "Book 50", "Author 2", 19.99},
	}
}

func filterBooksByAuthor(books []Book, filter string) (filtered []Book) {
	for _, book := range books {
		if strings.Contains(book.Author, filter) {
			filtered = append(filtered, book)
		}
	}

	return
}

func paginateBooks(books []Book, page, limit int) (paginated []Book) {
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	// Если начальный индекс выходит за границы, то возвращаем пустой список.
	if startIndex >= len(books) {
		paginated = []Book{}

		return
	}

	// Если конечный индекс превышает общее количество записей, то устанавливаем его в конец.
	if endIndex > len(books) {
		endIndex = len(books)
	}

	paginated = append(paginated, books[startIndex:endIndex]...)

	return
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("JSON encoding error: %v", err), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

// BooksHandler обрабатывает запросы к коллекции книг.
func BooksHandler(w http.ResponseWriter, r *http.Request) {
	filteredBooks := books

	r.ParseForm()

	// Извлекаем параметры запроса.
	pageParam := r.FormValue("page")
	limitParam := r.FormValue("limit")
	author := r.FormValue("author")

	if author != "" {
		// Инициируем фильтр книг на основе 'author'.
		filteredBooks = filterBooksByAuthor(filteredBooks, author)
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

		filteredBooks = paginateBooks(filteredBooks, page, limit)
	}

	// Отправляем подмножество книг в формате JSON.
	sendJSONResponse(w, filteredBooks)
}
