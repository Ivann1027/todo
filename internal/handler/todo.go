package handler

import (
	"encoding/json"
	"net/http"
	"todo-api/internal/model"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req model.CreateTodoDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	todo, err := h.services.Todo.CreateTodo(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.services.Todo.GetAllTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "todoId")
	if idStr == "" {
		http.Error(w, "Todo is is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.services.Todo.DeleteTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
