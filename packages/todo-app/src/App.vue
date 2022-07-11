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

fetchData();
</script>

<template>
  <TodoList :todos="todosData" @on-todo-added="(todo) => handleTodoAdded(todo)" />
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
