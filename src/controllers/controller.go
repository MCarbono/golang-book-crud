package controllers

import (
	"fmt"
	"go-bookstore/src/controllers/response"
	"go-bookstore/src/models"
	"go-bookstore/src/parse"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	response.Do(w, newBooks, http.StatusOK)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(id)
	response.Do(w, bookDetails, http.StatusOK)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	parse.UnmarshalJSON(r, createBook)
	b := createBook.CreateBook()
	response.Do(w, b, http.StatusCreated)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(id)
	response.Do(w, book, http.StatusNoContent)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	parse.UnmarshalJSON(r, updateBook)

	bookId := mux.Vars(r)["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookById(id)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	response.Do(w, bookDetails, http.StatusOK)
}
