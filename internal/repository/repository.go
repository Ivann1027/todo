package repository

import (
	"todo-api/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Todo interface {
	CreateTodo(todo *model.Todo) error
	GetAllTodos() ([]model.Todo, error)
	DeleteTodo(id uuid.UUID) error
}

type Repository struct {
	Todo
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Todo: NewTodoPostgres(db),
	}
}
