<script setup lang="ts">
import Todo from "./Todo.vue";
import type { Todo as TodoType } from "../types";
import { ref } from "vue";
import { createTodo } from "../api-client/todo";

defineProps<{ todos: TodoType[] }>();
const emit = defineEmits(["onTodoAdded"]);

const newTodoText = ref("");

async function handleNewTodoAdded() {
  const res = await createTodo({ title: newTodoText.value });
  emit("onTodoAdded", res);
  newTodoText.value = "";
}
</script>

<template>
  <div>
    <template v-for="(item, index) in todos"> <Todo v-bind="item" /></template>
    <div class="new-todo">
      <input type="text" v-model="newTodoText" />
      <button @click="handleNewTodoAdded">Add</button>
    </div>
  </div>
</template>

<style scoped></style>
