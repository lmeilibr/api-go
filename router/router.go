package router

import (
	"api-go/controller"
	"api-go/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func NewRouter(db db.AppDB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	r.Route("/api/users", func(r chi.Router) {
		r.Post("/", controller.HandleCreateUser(db))
		r.Get("/", controller.HandleGetAllUsers(db))
		r.Get("/{id}", controller.HandleGetUserByID(db))
		r.Put("/{id}", controller.HandleUpdateUser(db))
		r.Delete("/{id}", controller.HandleDeleteUser(db))
	})

	return r
}
