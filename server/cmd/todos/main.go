package main

import (
	"github.com/PieroNarciso/todo-hex/cmd/todos/controllers"
	"github.com/PieroNarciso/todo-hex/pkg/todos/infra"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	todoGroup := app.Group("/todos")

	todoRepository := infra.NewMemRepository()
	todoCreateController := controllers.NewCreateController(todoRepository)
	todoGroup.Post("/", todoCreateController.Exec) 

	todoListController := controllers.NewListController(todoRepository)
	todoGroup.Get("/", todoListController.Exec)

	todoGetController := controllers.NewGetController(todoRepository)
	todoGroup.Get("/:id", todoGetController.Exec)

	todoUpdateController := controllers.NewUpdateController(todoRepository)
	todoGroup.Put("/:id", todoUpdateController.Exec)


	app.Listen(":3000")
}
