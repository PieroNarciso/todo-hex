package domain

import "github.com/google/uuid"

type Todo struct {
	ID        uuid.UUID
	Title     string
	Completed bool
}

type TodoRepository interface {
	FindAll() ([]*Todo, error)
	FindByID(id uuid.UUID) (*Todo, error)
	UpdateOne(id uuid.UUID, todo *Todo) error
	DeleteOne(id uuid.UUID) error
	Save(todo *Todo) (*Todo, error)
}
