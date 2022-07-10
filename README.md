# todo-list

Todo List in Go and in Vue.

# Product Requirements

1. Can create a Todo. A todo consists of a title of any length
2. Todos can be updated - marked as completed.
3. Todos can be deleted.

# Technical Design

For the UI, it is developed as a Vue app.
For the API, Why REST? Wanted to learn how to make Go APIs in REST first prior to jumping to GraphQL, or grpc.

## Models

### Todo

```ts
interface Todo {
  id: number /* note, IRL this would be string but simplification  */;
  title: string;
  completed: boolean;
  created_date: Date;
  updated_date: Date;
}
```

## REST API

### GET /todos

Get all todos

### GET /todo/:id

Get one todo

### POST /todo

Create one todo

### PUT /todo/:id

Update Todo

### DELETE /todo/:id

Delete Todo
