<script setup lang="ts">
import Todo from "./Todo.vue";
import type { Todo as TodoType } from "../types";

defineProps<{ todos: TodoType[] }>();
</script>

<template>
  <template v-for="(item, index) in todos"> <Todo v-bind="item" /></template>
</template>

<style scoped>
.container {
  display: flex;
  align-items: center;
}
.title {
  display: block;
  width: 75%;
}
</style>
