# todo-list

Todo List in Go and in Vue

# Product Requirements

1. Can create a Todo. A todo consists of a title of any length
2. Todos can be updated - marked as completed.
3. Todos can be deleted.

# Models

## Todo

```ts
interface Todo {
  title: string;
  completed: boolean;
  created_date: Date;
  updated_date: Date;
}
```

# REST API

Why REST? Wanted to learn how to make Go APIs in REST first prior to jumping to GraphQL, or grpc.

## GET /todos

Get all todos

## GET /todo/:id

Get one todo

## POST /todo

Create one todo

## PUT /todo/:id

Update Todo

## DELETE /todo/:id

Delete Todo
