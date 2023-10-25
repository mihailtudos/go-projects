package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tudosm/go-projects/bookstore/pkg/models"
	"github.com/tudosm/go-projects/bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookId, _ := strconv.Atoi(params["bookId"])
	book, _ := models.GetBookById(bookId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(book)

	_, _ = w.Write(res)
}

func UpdateBooks(w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}
	utils.ParseBody(r, updatedBook)
	params := mux.Vars(r)
	bookId, _ := strconv.Atoi(params["bookId"])

	book, db := models.GetBookById(bookId)
	if updatedBook.Name != "" {
		book.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		book.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		book.Publication = updatedBook.Publication
	}

	db.Updates(book)
}

func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookId, _ := strconv.Atoi(params["bookId"])
	book := models.DeleteBook(bookId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(book)

	_, _ = w.Write(res)
}
