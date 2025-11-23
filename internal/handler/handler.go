package handler

import (
	"net/http"
	"todo-api/internal/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services}
}

func InitRouter(h *Handler) http.Handler {
	r := chi.NewRouter()

	// Health-check
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running!"))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/todos", h.CreateTodo)
		r.Get("/todos", h.GetAllTodos)
	})

	return r
}