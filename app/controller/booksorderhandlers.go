package controller

import (
	"bookexchange/app/models"
	"bookexchange/app/repository"
	"bookexchange/utils"
	"net/http"
)

type booksorder struct {
	db repository.BooksOrderDB
}

var t models.Tree

func NewBooksOrderRepository(pd repository.BooksOrderDB) repository.BooksOrder {
	return &booksorder{db: pd}
}

func (p *booksorder) AddBookOrder(w http.ResponseWriter, r *http.Request) {
	var data models.BookOrder

	if err := models.DecodeAndValidate(r, &data); err != nil {
		utils.WriteJsonErr(w, err.Error())
		return
	}

	tree := p.db.Insert(&t, data)

	utils.WriteJsonRes(w, tree)
	return
}
