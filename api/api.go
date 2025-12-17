package api

import (
	"net/http"
	"userCrud/models/application"
	"userCrud/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(db *application.Application) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/users", services.HandleSaveUser(db))
	r.Get("/api/users", services.HandleGetUsers(db))
	r.Get("/api/users/{id}", services.HandleGetUser(db))
	r.Put("/api/users/{id}", services.HandleUpdateUser(db))
	r.Delete("/api/users/{id}", services.HandleDeleteUser(db))

	return r
}
