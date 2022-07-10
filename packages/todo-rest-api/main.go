package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Todo struct {
	Id int
	Title string
	Completed bool

}


func listTodos(c *fiber.Ctx, db *sql.DB) error {
	var todos []Todo
	// return c.SendString("Hello, Worldo!")
	// query the db
	rows, err := db.Query("SELECT id, title, completed FROM todos ORDER BY id ASC")
	// close the rows, prevent further enumeration
	defer rows.Close()

	// TODO error handling is naive here
	if err != nil {
		log.Fatalln(err)
		return c.JSON("An error occurred")
	}

	for rows.Next() {
		todo := Todo{}
		rows.Scan(&todo.Id, &todo.Title, &todo.Completed)
		todos = append(todos, todo)
	}

	return c.JSON(todos)
}

func getTodoById(c *fiber.Ctx, db *sql.DB) error {
	todo := Todo{}

	row := db.QueryRow("SELECT id, title, completed FROM todos WHERE id = ($1)", c.Params("id"))

	err := row.Scan(&todo.Id, &todo.Title, &todo.Completed)

	if err != nil && err == sql.ErrNoRows {
		return c.SendStatus(404)
	}

	return c.JSON(todo)
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
	// TODO: environment variables
	// TODO: can we make this a singleton?
	connStr := "postgresql://user:postgrespassword@postgres:5432/todos?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/todos", func (c *fiber.Ctx) error {
		return listTodos(c, db)
	})
	app.Get("/todo/:id", func (c *fiber.Ctx) error {
		return getTodoById(c, db)
	})
	app.Post("/todo", createTodo)
	app.Put("/todo/:id", updateTodo)
	app.Delete("/todo/:id", deleteTodo)

	port := os.Getenv("PORT")
	if port == "" {
			port = "3000"
	}
	app.Listen(fmt.Sprintf(":%v", port))
}
