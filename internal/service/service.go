package service

import (
	"todo-api/internal/model"
	"todo-api/internal/repository"

	"github.com/google/uuid"
)

type Todo interface {
	CreateTodo(req model.CreateTodoDto) (*model.Todo, error)
	GetAllTodos() ([]model.Todo, error)
	DeleteTodo(id uuid.UUID) error
}

type Service struct {
	Todo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Todo: NewTodoService(repos.Todo),
	}
}
