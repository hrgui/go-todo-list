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

func createTodo(c *fiber.Ctx, db *sql.DB) error {
	newTodo := Todo{}

	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	fmt.Printf("%v", newTodo)

	row := db.QueryRow("INSERT into todos (title) VALUES ($1) RETURNING id, title, completed", newTodo.Title)
	row.Scan(&newTodo.Id, &newTodo.Title, &newTodo.Completed)

	return c.JSON(newTodo)
}

func updateTodo(c *fiber.Ctx, db *sql.DB) error {
	newTodo := Todo{}
	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	fmt.Printf("new todo: %v", newTodo)

	row := db.QueryRow("UPDATE todos SET title=($1), completed=($2) WHERE id=($3) RETURNING id, title, completed", newTodo.Title, newTodo.Completed, newTodo.Id)
	row.Scan(&newTodo.Id, &newTodo.Title, &newTodo.Completed)

	return c.JSON(newTodo)
}

func deleteTodo(c *fiber.Ctx, db *sql.DB) error {
	db.QueryRow("DELETE FROM todos WHERE id=($1)", c.Params("id"))

	return c.SendStatus(204)
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
	app.Post("/todo", func (c *fiber.Ctx) error {
		return createTodo(c, db)
	})
	app.Put("/todo/:id", func (c *fiber.Ctx) error {
		return updateTodo(c, db)
	})

	app.Delete("/todo/:id", func (c *fiber.Ctx) error {
		return deleteTodo(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
			port = "3000"
	}
	app.Listen(fmt.Sprintf(":%v", port))
}
