package controllers

import (
	"github.com/PieroNarciso/todo-hex/pkg/todos/domain"
	"github.com/gofiber/fiber/v2"
)


type todoListController struct {
	todoRepo domain.TodoRepository
}

func NewListController(todoRepository domain.TodoRepository) *todoListController {
	return &todoListController{
		todoRepo: todoRepository,
	}
}


func (c *todoListController) Exec(ctx *fiber.Ctx) error {
	todos, err := c.todoRepo.FindAll()
	if err != nil {
		return err
	}

	if todos == nil {
		todos = []*domain.Todo{}
	}
	return ctx.JSON(todos)
}
