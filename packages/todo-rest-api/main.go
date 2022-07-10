package main

import "github.com/gofiber/fiber/v2"

func listTodos(c *fiber.Ctx) error {
	return c.SendString("Hello, Worldo!")
}

func getTodoById(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func createTodo(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func updateTodo(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func deleteTodo(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}


func main() {
	app := fiber.New()

	app.Get("/todos", listTodos)
	app.Get("/todo/:id", getTodoById)
	app.Post("/todo", createTodo)
	app.Put("/todo/:id", updateTodo)
	app.Delete("/todo/:id", deleteTodo)

	app.Listen(":3000")
}
