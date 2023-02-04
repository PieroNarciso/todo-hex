package infra

import (
	"fmt"

	"github.com/PieroNarciso/todo-hex/pkg/todos/domain"
	"github.com/google/uuid"
)

type memRepository struct {
	todos []*domain.Todo
}

func NewMemRepository() *memRepository {
	return &memRepository{}
}

func (r *memRepository) FindAll() ([]*domain.Todo, error) {
	return r.todos, nil
}

func (r *memRepository) FindByID(id uuid.UUID) (*domain.Todo, error) {
	for _, todo := range r.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return nil, fmt.Errorf("todo not found")
}

func (r *memRepository) UpdateOne(id uuid.UUID, todo *domain.Todo) error {
	for i, t := range r.todos {
		if t.ID == id {
			r.todos[i] = todo
			return nil
		}
	}
	return fmt.Errorf("todo not found")
}

func (r *memRepository) DeleteOne(id uuid.UUID) error {
	for i, todo := range r.todos {
		if todo.ID == id {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo not found")
}

func (r *memRepository) Save(todo *domain.Todo) (*domain.Todo, error) {
	todo.ID = uuid.New()
	r.todos = append(r.todos, todo)
	return todo, nil
}
