# todo-list

Todo List in Go and in Vue.

# Product Requirements

1. Can create a Todo. A todo consists of a title of any length
2. Todos can be updated - marked as completed.
3. Todos can be deleted.

# Getting Started with Development

## You'll need / Prerequisites

- node.js v16 + yarn
- docker

## How to run - UI

```
yarn
```

```
yarn dev
```

## How to run - API

In root directory,

```
docker-compose up -d
```

## Ports

- 4000 - rest api

# Technical Design

For the UI, it is developed as a Vue app.
For the API, Why REST? Wanted to learn how to make Go APIs in REST first prior to jumping to GraphQL, or grpc.

The whole point is to learn how to go out of the comfort zone as a Software Engineer. I have done React + Node for too long. I am curious about Go and Vue, and I want to learn from it. I want to know what I can make from those technologies.

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

# How I learned all this / or reviewed

- https://docs.gofiber.io/
- https://medium.com/easyread/today-i-learned-golang-live-reload-for-development-using-docker-compose-air-ecc688ee076
- https://chartio.com/resources/tutorials/how-to-list-databases-and-tables-in-postgresql-using-psql/
- https://x-team.com/blog/automatic-timestamps-with-postgresql/
- https://blog.logrocket.com/building-simple-app-go-postgresql/ (this tutorial has some issues when the DB schema is just more than 1 column)
- https://go.dev/doc/database/querying
