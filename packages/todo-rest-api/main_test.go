package main

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func Test_deleteTodo(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := deleteTodo(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("deleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_updateTodo(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := updateTodo(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("updateTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createTodo(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createTodo(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("createTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getTodoById(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := getTodoById(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("getTodoById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_listTodos(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := listTodos(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("listTodos() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
