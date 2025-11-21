package repository

import (
	"todo-api/internal/model"

	"github.com/jackc/pgx/v5"
)

type TodoRepository interface {
	CreateTodo(todo *model.Todo) error
}

func NewRepository(db *pgx.Conn) TodoRepository {
	return &todoRepository{db}
}
