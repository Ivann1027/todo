package service

import (
	"todo-api/internal/model"
	"todo-api/internal/repository"
)

type Todo interface {
	CreateTodo(req model.CreateTodoDto) (*model.Todo, error)
	GetAllTodos() ([]model.Todo, error)
}

type Service struct {
	Todo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Todo: NewTodoService(repos.Todo),
	}
}
