package controllers

import (
	"encoding/json"

	"github.com/PieroNarciso/todo-hex/pkg/todos/domain"
	"github.com/gofiber/fiber/v2"
)

type createController struct {
	todoRepo domain.TodoRepository
}

type todoCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewCreateController(todoRepository domain.TodoRepository) *createController {
	return &createController{
		todoRepo: todoRepository,
	}
}

func (c *createController) Exec(ctx *fiber.Ctx) error {
	body := ctx.Body()

	var bodyRequest todoCreateRequest
	err := json.Unmarshal(body, &bodyRequest)

	if err != nil {
		return err
	}

	todo, err := c.todoRepo.Save(&domain.Todo{
		Title:       bodyRequest.Title,
		Description: bodyRequest.Description,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(todo)
}
