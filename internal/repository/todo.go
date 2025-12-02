package repository

import (
	"context"
	"errors"
	"fmt"
	"todo-api/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TodoPostgres struct {
	db *pgx.Conn
}

func NewTodoPostgres(db *pgx.Conn) *TodoPostgres {
	return &TodoPostgres{db}
}

func (r *TodoPostgres) CreateTodo(todo *model.Todo) error {
	query := "insert into todos (id, text, is_done) values ($1, $2, $3);"

	_, err := r.db.Exec(context.Background(), query, todo.Id, todo.Text, todo.IsDone)
	if err != nil {
		return fmt.Errorf("failed to create todo: %w", err)
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

func (r *TodoPostgres) DeleteTodo(id uuid.UUID) error {
	query := "delete from todos where id = $1;"

	result, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("todo with id %s not found", id)
	}

	return nil
}

func (r *TodoPostgres) UpdateTodo(id uuid.UUID, dto model.UpdateTodoDto) (model.Todo, error) {
	textParam := any(nil)
	if dto.Text != nil {
		textParam = *dto.Text
	}

	isDoneParam := any(nil)
	if dto.IsDone != nil {
		isDoneParam = *dto.IsDone
	}

	query := "update todos set text = coalesce($2, text), is_done = coalesce($3, is_done) where id = $1 returning id, text, is_done"

	var todo model.Todo

	err := r.db.QueryRow(context.Background(), query, id, textParam, isDoneParam).
		Scan(&todo.Id, &todo.Text, &todo.IsDone)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Todo{}, fmt.Errorf("todo with id %s not found", id)
		}
		return model.Todo{}, fmt.Errorf("failed to update todo: %w", err)
	}

	return todo, nil
}
