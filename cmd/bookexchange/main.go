package main

import (
	"bookexchange/app/controller"
	"bookexchange/app/routes"
	"bookexchange/app/services"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

func main() {
	db := services.NewRedisConnection("redis://localhost:6379")
	r := chi.NewRouter()
	// Basic CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	DB := services.NewDbClient(db)
	ctrl := controller.NewBooksOrderRepository(DB)

	routes.NewRouter(ctrl, r)

	fmt.Println("server started on PORT 8070")
	http.ListenAndServe(":8070", r)
}