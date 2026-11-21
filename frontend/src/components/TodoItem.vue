<script setup>
import { nextTick, ref } from 'vue'

const props = defineProps({
  subtask: {
    type: Object,
    required: true,
  },
})

const emit = defineEmits(['update', 'remove'])

const isEditing = ref(false)
const draft = ref('')
const inputRef = ref(null)

async function startEditing() {
  draft.value = props.subtask.title
  isEditing.value = true

  await nextTick()
  inputRef.value?.focus()
  inputRef.value?.select()
}

function save() {
  const title = draft.value.trim()
  if (!title) {
    inputRef.value?.focus()
    return
  }

  if (title === props.subtask.title) {
    isEditing.value = false
    return
  }

  emit('update', title)
  isEditing.value = false
}

function cancel() {
  isEditing.value = false
}
</script>

<template>
  <div class="todo-item">
    <span class="todo-bullet">
      <i class="bi bi-check2"></i>
    </span>

    <template v-if="isEditing">
      <input
        ref="inputRef"
        v-model="draft"
        class="form-control todo-edit-input"
        type="text"
        maxlength="200"
        aria-label="Название подзадачи"
        @keyup.enter="save"
        @keyup.esc="cancel"
      />
      <div class="todo-actions">
        <button class="mini-btn" type="button" title="Сохранить" aria-label="Сохранить подзадачу" @click="save">
          <i class="bi bi-check-lg"></i>
        </button>
        <button class="mini-btn" type="button" title="Отмена" aria-label="Отменить редактирование" @click="cancel">
          <i class="bi bi-x-lg"></i>
        </button>
      </div>
    </template>

    <template v-else>
      <span class="todo-text">{{ subtask.title }}</span>
      <div class="todo-actions">
        <button
          class="mini-btn"
          type="button"
          title="Редактировать подзадачу"
          aria-label="Редактировать подзадачу"
          @click="startEditing"
        >
          <i class="bi bi-pencil"></i>
        </button>
        <button
          class="mini-btn danger"
          type="button"
          title="Удалить подзадачу"
          aria-label="Удалить подзадачу"
          @click="emit('remove')"
        >
          <i class="bi bi-trash3"></i>
        </button>
      </div>
    </template>
  </div>
</template>
