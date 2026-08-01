package service

import (
	"testing"
	"todo-api/internal/model"

	"github.com/google/uuid"
)

// Repository mocks
type MockTodoRepository struct {
	todos map[uuid.UUID]model.Todo
}
func (m *MockTodoRepository) CreateTodo(todo *model.Todo) error {
	m.todos[todo.Id] = *todo
	return nil
}
func (m *MockTodoRepository) GetAllTodos() ([]model.Todo, error) {
	return nil, nil
}
func (m *MockTodoRepository) DeleteTodo(id uuid.UUID) error {
	return nil
}
func (m *MockTodoRepository) UpdateTodo(id uuid.UUID, dto model.UpdateTodoDto) (model.Todo, error) {
	return model.Todo{}, nil
}

// Test functions
func TestCreateTodo(t *testing.T) {
	mockRepo := &MockTodoRepository{
		todos: make(map[uuid.UUID]model.Todo),
	}
	todoService := NewTodoService(mockRepo)

	tests := []struct{
		name string
		reqData model.CreateTodoDto
		wantError bool
	}{
		{"empty text", model.CreateTodoDto{Text: ""}, true},
		{"normal text", model.CreateTodoDto{Text: "Починить машину"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := todoService.CreateTodo(tt.reqData)
			if tt.wantError {
				 if err == nil {
					t.Errorf("CreateTodo expected error, got %+v", got)
				 }
				 return
			}
			if err != nil {
				t.Errorf("CreateTodo unexpected error: %s", err)
			}
			if got.Id == uuid.Nil {
				t.Errorf("Expected valid UUID, got nil")
			}
			if got.Text != tt.reqData.Text {
				t.Errorf("Expected text - %s, got - %s", tt.reqData.Text, got.Text)
			}
			if got.IsDone != false {
				t.Errorf("Expected isDone=false, got %v", got.IsDone)
			}
		})
	}
}