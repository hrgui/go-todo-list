import { Todo } from "../types";

export async function createTodo(todo: Partial<Todo>): Promise<Todo> {
  const res = await fetch(`/api/todo`, {
    method: "post",
    body: JSON.stringify(todo),
    headers: {
      "Content-Type": "application/json",
    },
  });
  return res.json();
}

export async function updateTodo(todo: Partial<Todo>): Promise<Todo> {
  const res = await fetch(`/api/todo/${todo.id}`, {
    method: "put",
    body: JSON.stringify(todo),
    headers: {
      "Content-Type": "application/json",
    },
  });
  return res.json();
}

export async function listTodos(): Promise<Todo[]> {
  const todos = await fetch(`/api/todos`);
  return todos.json();
}

export function deleteTodo(id: number) {
  return fetch(`/api/todo/${id}`, { method: "delete" });
}
