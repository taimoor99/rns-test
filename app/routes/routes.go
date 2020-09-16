package routes

import (
	"bookexchange/app/repository"
	"github.com/go-chi/chi"
)

func NewRouter(ctrl repository.BooksOrder, r *chi.Mux) {

	r.Group(func(r chi.Router) {
		r.Post("/addorder", ctrl.AddBookOrder)
	})

	return
}

