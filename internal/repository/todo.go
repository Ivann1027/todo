package repository

import (
	"context"
	"fmt"
	"todo-api/internal/model"

	"github.com/jackc/pgx/v5"
)

type TodoPostgres struct {
	db *pgx.Conn
}

func NewTodoPostgres(db *pgx.Conn) *TodoPostgres {
	return &TodoPostgres{db}
}

func (r *TodoPostgres) CreateTodo(todo *model.Todo) error {
	query := "insert into todos (id, text, is_done) values ($1, $2, $3) returning id;"

	err := r.db.QueryRow(context.Background(), query, todo.Id, todo.Text, todo.IsDone).Scan(&todo.Id)
	if err != nil {
		return fmt.Errorf("failed to create todo in db: %w", err)
	}

	return nil
}

func (r *TodoPostgres) GetAllTodos() ([]model.Todo, error) {
	var todos []model.Todo
	query := "select * from todos;"
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all todos in db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.Id, &todo.Text, &todo.IsDone); err != nil {
			return nil, fmt.Errorf("failed to scan todo: %w", err)
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return todos, nil
}
