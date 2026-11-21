import { API_ENDPOINTS } from '../client/api-endpoints'
import { apiClient } from '../client/api-client'
import { errorProcessing } from '../helpers/error-processing'
import { normalizeSmallTask } from './small-task.api'

// Backend отдаёт ID/CreatedAt/UpdatedAt в верхнем регистре, а title/description — в нижнем.
const normalizeTask = (task) => ({
  id: task.id ?? task.ID,
  title: task.title ?? '',
  description: task.description ?? '',
  createdAt: task.createdAt ?? task.CreatedAt,
  updatedAt: task.updatedAt ?? task.UpdatedAt,
  smallTasks: Array.isArray(task.small_tasks)
    ? task.small_tasks.map(normalizeSmallTask)
    : [],
})

class TaskApi {
  async getTasks() {
    try {
      const tasks = await apiClient.get(API_ENDPOINTS.TASKS.LIST)
      return Array.isArray(tasks) ? tasks.map(normalizeTask) : []
    } catch (err) {
      throw errorProcessing(err, 'получении задач')
    }
  }

  async createTask(data) {
    try {
      const task = await apiClient.post(API_ENDPOINTS.TASKS.CREATE, data)
      return normalizeTask(task)
    } catch (err) {
      throw errorProcessing(err, 'создании задачи')
    }
  }

  async updateTask(id, data) {
    try {
      const task = await apiClient.patch(API_ENDPOINTS.TASKS.UPDATE(id), data)
      return normalizeTask(task)
    } catch (err) {
      throw errorProcessing(err, 'обновлении задачи')
    }
  }

  async deleteTask(id) {
    try {
      return await apiClient.delete(API_ENDPOINTS.TASKS.DELETE(id))
    } catch (err) {
      throw errorProcessing(err, 'удалении задачи')
    }
  }
}

export const taskApi = new TaskApi()
