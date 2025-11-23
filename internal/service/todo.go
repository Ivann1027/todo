package service

import (
	"errors"
	"fmt"
	"todo-api/internal/model"
	"todo-api/internal/repository"

	"github.com/google/uuid"
)

type TodoService struct {
	todoRepo repository.Todo
}

func NewTodoService(todoRepo repository.Todo) *TodoService {
	return &TodoService{todoRepo}
}

func (s *TodoService) CreateTodo(req model.CreateTodoDto) (*model.Todo, error) {
	if req.Text == "" {
		return nil, errors.New("text cannot be empty")
	}

	if len(req.Text) > 500 {
		return nil, errors.New("text is too long (max 500 characters)")
	}

	todo := &model.Todo{
		Id:     uuid.New(),
		Text:   req.Text,
		IsDone: false,
	}

	err := s.todoRepo.CreateTodo(todo)
	if err != nil {
		return nil, fmt.Errorf("failed to create todo in service: %w", err)
	}

	return todo, nil
}

func (s *TodoService) GetAllTodos() ([]model.Todo, error) {
	todos, err := s.todoRepo.GetAllTodos()
	if err != nil {
		return nil, err
	}

	return todos, nil
}
