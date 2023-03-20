package controllers

import (
	"github.com/PieroNarciso/todo-hex/pkg/todos/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type todoUpdateController struct {
	todoRepository domain.TodoRepository
}

func NewUpdateController(todoRepository domain.TodoRepository) *todoUpdateController {
	return &todoUpdateController{
		todoRepository: todoRepository,
	}
}

type todoUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   *bool  `json:"completed"`
}

func (c *todoUpdateController) Exec(ctx *fiber.Ctx) error {
	var updateRequest todoUpdateRequest

	err := ctx.BodyParser(&updateRequest)
	if err != nil {
		return err
	}
	todoId := ctx.Params("id")
	todoUuid, err := uuid.Parse(todoId)
	if err != nil {
		return err
	}
	todo, err := c.todoRepository.FindByID(todoUuid)
	if err != nil {
		return err
	}

	if updateRequest.Title != "" {
		todo.Title = updateRequest.Title
	}
	if updateRequest.Completed != nil {
		todo.Completed = *updateRequest.Completed
	}
	if updateRequest.Description != "" {
		todo.Description = updateRequest.Description
	}

	err = c.todoRepository.UpdateOne(todoUuid, todo)
	if err != nil {
		return err
	}
	return ctx.JSON(todo)
}
