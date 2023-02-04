package infra_test

import (
	"testing"

	"github.com/PieroNarciso/todo-hex/pkg/todos/domain"
	"github.com/PieroNarciso/todo-hex/pkg/todos/infra"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {

	t.Run("should return none users", func(t *testing.T) {
		repo := infra.NewMemRepository()

		todos, _ := repo.FindAll()
		assert.Equal(t, 0, len(todos))
	})

	t.Run("should return one user", func(t *testing.T) {
		repo := infra.NewMemRepository()
		_, err := repo.Save(&domain.Todo{
			Title:     "test",
			Completed: false,
		})
		assert.NoError(t, err)

		todos, _ := repo.FindAll()
		assert.Equal(t, 1, len(todos))
	})
}

func TestFindById(t *testing.T) {
	t.Run("should return one user", func(t *testing.T) {
		repo := infra.NewMemRepository()
		todoCreated, err := repo.Save(&domain.Todo{
			Title:     "test",
			Completed: false,
		})
		assert.NoError(t, err)

		todo, err := repo.FindByID(todoCreated.ID)
		assert.NoError(t, err)
		assert.Equal(t, "test", todo.Title)
	})

	t.Run("should return error", func(t *testing.T) {
		repo := infra.NewMemRepository()
		_, err := repo.Save(&domain.Todo{
			Title:     "test",
			Completed: false,
		})
		assert.NoError(t, err)

		_, err = repo.FindByID(uuid.New())
		assert.Error(t, err)
	})
}

func TestUpdateOne(t *testing.T) {
	t.Run("should update a user", func(t *testing.T) {
		repo := infra.NewMemRepository()
		todoCreated, err := repo.Save(&domain.Todo{
			Title:     "test",
			Completed: false,
		})
		assert.NoError(t, err)

		todo, _ := repo.FindByID(todoCreated.ID)
		todo.Title = "test2"
		repo.UpdateOne(todo.ID, todo)

		todo, _ = repo.FindByID(todoCreated.ID)
		assert.Equal(t, "test2", todo.Title)
	})

	t.Run("should return error", func(t *testing.T) {
		repo := infra.NewMemRepository()
		todoCreated, err := repo.Save(&domain.Todo{
			Title:     "test",
			Completed: false,
		})
		assert.NoError(t, err)

		todo, _ := repo.FindByID(todoCreated.ID)
		todo.Title = "test2"
		err = repo.UpdateOne(uuid.New(), todo)
		assert.Error(t, err)
	})
}

func TestDeleteOne(t *testing.T) {
	t.Run("should delete a user", func(t *testing.T) {
		repo := infra.NewMemRepository()
		todoCreated, err := repo.Save(&domain.Todo{
			Title:     "test",
			Completed: false,
		})
		assert.NoError(t, err)

		repo.DeleteOne(todoCreated.ID)
		todos, _ := repo.FindAll()
		assert.Equal(t, 0, len(todos))
	})

	t.Run("should return error", func(t *testing.T) {
		repo := infra.NewMemRepository()
		_, err := repo.Save(&domain.Todo{
			Title:     "test",
			Completed: false,
		})
		assert.NoError(t, err)

		err = repo.DeleteOne(uuid.New())
		assert.Error(t, err)

		todos, _ := repo.FindAll()
		assert.Equal(t, 1, len(todos))
	})
}

func TestSave(t *testing.T) {
	t.Run("should save a user", func(t *testing.T) {
		repo := infra.NewMemRepository()
		_, err := repo.Save(&domain.Todo{
			Title:     "test",
			Completed: false,
		})
		assert.NoError(t, err)

		todos, err := repo.FindAll()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(todos))
	})
}
