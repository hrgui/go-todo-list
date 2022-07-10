<script setup lang="ts">
import { ref } from "vue";
import TodoList from "./components/TodoList.vue";
import { Todo as TodoType } from "./types";

const todosData = ref<TodoType[]>([]);

async function fetchData() {
  todosData.value = [];
  const res = await fetch(`https://jsonplaceholder.typicode.com/todos`);
  todosData.value = await res.json();
}

fetchData();
</script>

<template>
  <TodoList :todos="todosData" />
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
