package repository

import (
	"context"
	"fmt"
	"todo-api/internal/model"

	"github.com/jackc/pgx/v5"
)

type todoRepository struct {
	db *pgx.Conn
}

func (r *todoRepository) CreateTodo(todo *model.Todo) error {
	query := "insert into todos (id, text, is_done) values ($1, $2, $3) returning id;"

	err := r.db.QueryRow(context.Background(), query, todo.Id, todo.Text, todo.IsDone).Scan(&todo.Id)
	if err != nil {
		return fmt.Errorf("failed to create todo in db: %w", err)
	}

	return nil
}
