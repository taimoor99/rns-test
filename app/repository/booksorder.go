package repository

import (
	"bookexchange/app/models"
	"net/http"
)

type BooksOrder interface{
	AddBookOrder(w http.ResponseWriter, r *http.Request)
}

type BooksOrderDB interface{
	AddBook(order models.BookOrder) error
	RemoveBook() error
	Insert(t *models.Tree, v models.BookOrder) *models.Tree
}