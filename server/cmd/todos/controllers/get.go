package controllers

import (
	"net/http"

	"github.com/PieroNarciso/todo-hex/pkg/todos/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type todoGetController struct {
	todoRepository domain.TodoRepository
}

func NewGetController(todoRepository domain.TodoRepository) *todoGetController {
	return &todoGetController{
		todoRepository: todoRepository,
	}
}

func (c *todoGetController) Exec(ctx *fiber.Ctx) error {
	todoId := ctx.Params("id")

	uuidTodo, err := uuid.Parse(todoId)
	if err != nil {
		return err
	}

	todo, err := c.todoRepository.FindByID(uuidTodo)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(todo)
}
