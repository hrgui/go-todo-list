<script setup lang="ts">
import { ref } from "vue";
import { listTodos } from "./api-client/todo";
import TodoList from "./components/TodoList.vue";
import { Todo, Todo as TodoType } from "./types";

const todosData = ref<TodoType[]>([]);

async function fetchData() {
  todosData.value = [];
  todosData.value = await listTodos();
}

function handleTodoAdded(todo: Todo) {
  todosData.value.push(todo);
}

function handleTodoDeleted(todo: Todo) {
  todosData.value = todosData.value.filter((t) => t.id !== todo.id);
}

fetchData();
</script>

<template>
  <TodoList
    :todos="todosData"
    @on-todo-added="(todo) => handleTodoAdded(todo)"
    @on-todo-deleted="(todo) => handleTodoDeleted(todo)"
  />
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
