<script setup lang="ts">
import { ref, watch } from "vue";
import { deleteTodo, updateTodo } from "../api-client/todo";

const emit = defineEmits(["onTodoDeleted"]);
const props = defineProps<{ title: string; id: number; completed: boolean }>();
const isEditing = ref(false);
const isCompleted = ref(props.completed);

const todoNonEditClass = ref(isCompleted.value ? "title-completed" : "title-non-edit");

async function handleToggleCompleted() {
  await updateTodo({ title: props.title, completed: isCompleted.value, id: props.id });
  todoNonEditClass.value = isCompleted.value ? "title-completed" : "title-non-edit";
}

watch(isCompleted, handleToggleCompleted);

async function handleSave() {
  await updateTodo({ title: props.title, completed: props.completed, id: props.id });
  isEditing.value = false;
}

async function handleDelete() {
  await deleteTodo(props.id);
  isEditing.value = false;
  emit("onTodoDeleted", props);
}
</script>

<template>
  <div class="container">
    <div class="title">
      <input type="text" v-model="title" v-if="isEditing" />
      <button v-if="isEditing" @click="handleSave">Save</button>
      <button v-if="isEditing" @click="handleDelete">Delete</button>
      <button
        :title="`Edit ${title}`"
        :class="todoNonEditClass"
        @click="isEditing = !isEditing"
        v-if="!isEditing"
      >
        {{ title }}
      </button>
    </div>
    <input type="checkbox" v-model="isCompleted" />
  </div>
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

.title-non-edit {
  background: none;
  color: inherit;
  border: none;
  padding: 0;
  font: inherit;
  cursor: pointer;
  outline: inherit;
}

.title-completed {
  background: none;
  color: inherit;
  border: none;
  padding: 0;
  font: inherit;
  cursor: pointer;
  outline: inherit;
  text-decoration: line-through;
}
</style>
