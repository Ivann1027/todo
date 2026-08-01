package model

import "github.com/google/uuid"

type Todo struct {
	Id     uuid.UUID `json:"id"`
	Text   string    `json:"text"`
	IsDone bool      `json:"isDone"`
}

type CreateTodoDto struct {
	Text string `json:"text"`
}

type UpdateTodoDto struct {
	Text   *string `json:"text,omitempty"`
	IsDone *bool   `json:"isDone,omitempty"`
}
