package domain

import "github.com/google/uuid"

type Todo struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
}

type TodoRepository interface {
	FindAll() ([]*Todo, error)
	FindByID(id uuid.UUID) (*Todo, error)
	UpdateOne(id uuid.UUID, todo *Todo) error
	DeleteOne(id uuid.UUID) error
	Save(todo *Todo) (*Todo, error)
}
