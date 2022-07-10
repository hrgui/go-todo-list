<script setup lang="ts">
import { ref, watch } from "vue";

const props = defineProps<{ title: string; id: number; completed: boolean }>();
const isEditing = ref(false);
const isCompleted = ref(props.completed);

const todoNonEditClass = ref(isCompleted.value ? "title-completed" : "title-non-edit");

function onToggleCompleted() {
  //TODO: Need to run the update (save)
  todoNonEditClass.value = isCompleted.value ? "title-completed" : "title-non-edit";
}

watch(isCompleted, onToggleCompleted);

function onSave() {
  alert(`TODO: todo will be saved with ${props.title} ${props.id}`);
  isEditing.value = false;
}

function onDelete() {
  alert(`TODO: todo will be deleted ${props.id}`);
  isEditing.value = false;
}
</script>

<template>
  <div class="container">
    <div class="title">
      <input type="text" v-model="title" v-if="isEditing" />
      <button v-if="isEditing" @click="onSave">Save</button>
      <button v-if="isEditing" @click="onDelete">Delete</button>
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
