package service

import (
	"todo-api/internal/model"
	"todo-api/internal/repository"
)

type TodoService interface {
	CreateTodo(req model.CreateTodoDto) (*model.Todo, error)
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{todoRepo}
}
