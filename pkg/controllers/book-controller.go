package controllers

import (
	"encoding/json"
	"github.com/VeRJiL/go-bookstore/pkg/models"
	"github.com/VeRJiL/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var bookModel models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()

	res, _ := json.Marshal(books)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		panic(err)
	}

	book, _ := models.GetBookById(bookId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(book)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	bookModel := &models.Book{}
	utils.ParseBody(r, bookModel)
	book := bookModel.Create()
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		panic(err)
	}

	book := models.Delete(bookId)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//func updateBook(w http.ResponseWriter, r *http.Request) {
//	var updateBook := models.book{}
//	vars := mux.Vars(r)
//	bookId, err := strconv.ParseInt(vars["bookId"])
//	if err !+nil {
//		panic(err)
//	}
//
//	book, db := models.GetBookById(bookId)
//	if updateBook.name != "" {
//		updateBook.name = boo
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write(res)
//}
