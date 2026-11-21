<script setup>
import { computed, onMounted, ref } from 'vue'
import Swal from 'sweetalert2'

import TodoItem from '../components/TodoItem.vue'
import { taskApi } from '../api/services/task.api'
import { smallTaskApi } from '../api/services/small-task.api'

const tasks = ref([])
const selectedId = ref(null)
const isLoading = ref(true)

const subtasks = ref([])
const isSubtasksLoading = ref(false)
const newSubtaskTitle = ref('')
const subtaskRequestId = ref(0)

const selectedTask = computed(
  () => tasks.value.find((task) => task.id === selectedId.value) ?? null,
)

function showError(error) {
  Swal.fire({
    icon: 'error',
    title: 'Ошибка',
    text: error?.message || 'Не удалось выполнить запрос',
  })
}

async function loadSubtasks(taskId) {
  const requestId = ++subtaskRequestId.value
  subtasks.value = []
  isSubtasksLoading.value = true

  try {
    const items = await smallTaskApi.getSmallTasks(taskId)
    if (requestId !== subtaskRequestId.value) return
    subtasks.value = items
  } catch (error) {
    if (requestId !== subtaskRequestId.value) return
    showError(error)
  } finally {
    if (requestId === subtaskRequestId.value) {
      isSubtasksLoading.value = false
    }
  }
}

function selectTask(taskId) {
  selectedId.value = taskId
  newSubtaskTitle.value = ''
  loadSubtasks(taskId)
}

async function loadTasks() {
  isLoading.value = true

  try {
    tasks.value = await taskApi.getTasks()

    if (tasks.value.length) {
      selectTask(tasks.value[0].id)
    } else {
      selectedId.value = null
      subtasks.value = []
    }
  } catch (error) {
    showError(error)
  } finally {
    isLoading.value = false
  }
}

function promptTaskData(task = null) {
  return Swal.fire({
    title: task ? 'Редактировать задачу' : 'Новая задача',
    html: `
      <label class="swal2-input-label" for="taskTitle">Название</label>
      <input id="taskTitle" class="swal2-input" type="text" maxlength="200" placeholder="Например, Купить продукты" />
      <label class="swal2-input-label" for="taskDescription">Описание</label>
      <textarea id="taskDescription" class="swal2-input swal2-textarea" maxlength="1000" placeholder="Детали задачи"></textarea>
    `,
    focusConfirm: false,
    showCancelButton: true,
    confirmButtonText: task ? 'Сохранить' : 'Создать',
    cancelButtonText: 'Отмена',
    reverseButtons: true,
    didOpen: () => {
      const titleInput = document.getElementById('taskTitle')
      const descriptionInput = document.getElementById('taskDescription')

      titleInput.value = task?.title ?? ''
      descriptionInput.value = task?.description ?? ''
      titleInput.focus()
    },
    preConfirm: () => {
      const title = document.getElementById('taskTitle').value.trim()
      const description = document.getElementById('taskDescription').value.trim()

      if (!title) {
        Swal.showValidationMessage('Введите название задачи')
        return false
      }

      return { title, description }
    },
  })
}

async function createTask() {
  const { value } = await promptTaskData()
  if (!value) return

  try {
    const task = await taskApi.createTask(value)
    tasks.value.push(task)
    selectTask(task.id)
  } catch (error) {
    showError(error)
  }
}

async function editTask() {
  const task = selectedTask.value
  if (!task) return

  const { value } = await promptTaskData(task)
  if (!value) return

  try {
    const updatedTask = await taskApi.updateTask(task.id, value)
    const index = tasks.value.findIndex((item) => item.id === task.id)

    if (index !== -1) {
      tasks.value[index] = updatedTask
    }
  } catch (error) {
    showError(error)
  }
}

async function deleteTask() {
  const task = selectedTask.value
  if (!task) return

  const result = await Swal.fire({
    title: 'Удалить задачу?',
    text: `«${task.title}» будет удалена.`,
    icon: 'warning',
    showCancelButton: true,
    confirmButtonText: 'Удалить',
    cancelButtonText: 'Отмена',
    reverseButtons: true,
  })

  if (!result.isConfirmed) return

  try {
    await taskApi.deleteTask(task.id)
    tasks.value = tasks.value.filter((item) => item.id !== task.id)

    if (tasks.value.length) {
      selectTask(tasks.value[0].id)
    } else {
      selectedId.value = null
      subtasks.value = []
      newSubtaskTitle.value = ''
      subtaskRequestId.value += 1
    }
  } catch (error) {
    showError(error)
  }
}

async function addSubtask() {
  const task = selectedTask.value
  const title = newSubtaskTitle.value.trim()
  if (!task || !title) return

  try {
    const subtask = await smallTaskApi.createSmallTask({ taskId: task.id, title })

    if (selectedId.value === task.id) {
      subtasks.value.push(subtask)
    }

    newSubtaskTitle.value = ''
  } catch (error) {
    showError(error)
  }
}

async function updateSubtask(subtask, title) {
  const task = selectedTask.value
  if (!task) return

  try {
    const updated = await smallTaskApi.updateSmallTask(subtask.id, {
      taskId: task.id,
      title,
    })
    const index = subtasks.value.findIndex((item) => item.id === subtask.id)

    if (index !== -1) {
      subtasks.value[index] = updated
    }
  } catch (error) {
    showError(error)
  }
}

async function removeSubtask(subtask) {
  const result = await Swal.fire({
    title: 'Удалить подзадачу?',
    text: `«${subtask.title}» будет удалена.`,
    icon: 'warning',
    showCancelButton: true,
    confirmButtonText: 'Удалить',
    cancelButtonText: 'Отмена',
    reverseButtons: true,
  })

  if (!result.isConfirmed) return

  try {
    await smallTaskApi.deleteSmallTask(subtask.id)
    subtasks.value = subtasks.value.filter((item) => item.id !== subtask.id)
  } catch (error) {
    showError(error)
  }
}

onMounted(loadTasks)
</script>

<template>
  <div class="page-heading">
    <div>
      <p class="eyebrow">Список дел</p>
      <h1>Мои задачи</h1>
      <p class="page-subtitle">
        Задачи и подзадачи сохраняются на сервере: создавайте, редактируйте и удаляйте их.
      </p>
    </div>

    <button class="btn btn-highlight" type="button" @click="createTask">
      <i class="bi bi-plus-lg me-1"></i>
      Новая задача
    </button>
  </div>

  <div class="row g-4">
    <div class="col-lg-4 col-xl-3">
      <aside class="panel notes-panel">
        <div class="panel-head">
          <h2 class="panel-title">Задачи</h2>
        </div>

        <nav v-if="tasks.length" class="notes-list" aria-label="Список задач">
          <button
            v-for="task in tasks"
            :key="task.id"
            class="note-item"
            :class="{ active: task.id === selectedId }"
            type="button"
            @click="selectTask(task.id)"
          >
            <span class="note-item-icon">
              <i class="bi bi-check2-square"></i>
            </span>
            <span class="note-item-title">{{ task.title }}</span>
            <i class="bi bi-chevron-right note-chevron"></i>
          </button>
        </nav>

        <div v-else class="notes-empty">
          <i class="bi" :class="isLoading ? 'bi-hourglass-split' : 'bi-inbox'"></i>
          <p>{{ isLoading ? 'Загружаем задачи...' : 'Пока нет ни одной задачи.' }}</p>
        </div>
      </aside>
    </div>

    <div class="col-lg-8 col-xl-9">
      <section class="panel editor-panel">
        <template v-if="selectedTask">
          <header class="editor-head">
            <div class="editor-heading">
              <p class="eyebrow mb-1">Задача</p>
              <h2 class="editor-title">{{ selectedTask.title }}</h2>
            </div>

            <div class="editor-actions">
              <button
                class="icon-btn"
                type="button"
                title="Редактировать задачу"
                aria-label="Редактировать задачу"
                @click="editTask"
              >
                <i class="bi bi-pencil"></i>
              </button>
              <button
                class="icon-btn icon-btn-danger"
                type="button"
                title="Удалить задачу"
                aria-label="Удалить задачу"
                @click="deleteTask"
              >
                <i class="bi bi-trash3"></i>
              </button>
            </div>
          </header>

          <p
            class="task-description"
            :class="{ 'task-description-empty': !selectedTask.description }"
          >
            {{ selectedTask.description || 'Описание не указано.' }}
          </p>

          <section class="subtasks">
            <div class="subtasks-head">
              <h3 class="subtasks-title">Подзадачи</h3>
              <span v-if="subtasks.length" class="subtasks-count">{{ subtasks.length }}</span>
            </div>

            <div v-if="isSubtasksLoading" class="subtasks-empty">
              <i class="bi bi-hourglass-split"></i>
              <p>Загружаем подзадачи...</p>
            </div>

            <div v-else-if="subtasks.length" class="todo-list">
              <TodoItem
                v-for="subtask in subtasks"
                :key="subtask.id"
                :subtask="subtask"
                @update="updateSubtask(subtask, $event)"
                @remove="removeSubtask(subtask)"
              />
            </div>

            <div v-else class="subtasks-empty">
              <i class="bi bi-list-check"></i>
              <p>Пока нет подзадач. Добавьте первую.</p>
            </div>

            <form class="add-task" @submit.prevent="addSubtask">
              <input
                v-model="newSubtaskTitle"
                class="form-control"
                type="text"
                maxlength="200"
                placeholder="Новая подзадача"
                aria-label="Название подзадачи"
              />
              <button
                class="btn btn-accent"
                type="submit"
                :disabled="!newSubtaskTitle.trim()"
              >
                <i class="bi bi-plus-lg me-1"></i>
                Добавить
              </button>
            </form>
          </section>
        </template>

        <div v-else class="editor-empty">
          <i class="bi bi-journal-check"></i>
          <h2>Задача не выбрана</h2>
          <p>Выберите задачу из списка или создайте новую.</p>
        </div>
      </section>
    </div>
  </div>
</template>
