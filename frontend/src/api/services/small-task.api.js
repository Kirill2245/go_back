import { API_ENDPOINTS } from '../client/api-endpoints'
import { apiClient } from '../client/api-client'
import { errorProcessing } from '../helpers/error-processing'

// Backend может отдавать поля как в camelCase, так и в верхнем регистре.
export const normalizeSmallTask = (smallTask) => ({
  id: smallTask.id ?? smallTask.ID,
  taskId: smallTask.taskId ?? smallTask.task_id ?? smallTask.TaskID ?? smallTask.TaskId,
  title: smallTask.title ?? '',
  createdAt: smallTask.createdAt ?? smallTask.CreatedAt,
  updatedAt: smallTask.updatedAt ?? smallTask.UpdatedAt,
})

class SmallTaskApi {
  async getSmallTasks(taskID) {
    try {
      const smallTasks = await apiClient.get(API_ENDPOINTS.SMALL_TASKS.LIST(taskID))
      return Array.isArray(smallTasks) ? smallTasks.map(normalizeSmallTask) : []
    } catch (err) {
      throw errorProcessing(err, 'получении подзадач')
    }
  }

  async createSmallTask({ taskId, title }) {
    try {
      const smallTask = await apiClient.post(API_ENDPOINTS.SMALL_TASKS.CREATE, {
        task_id: taskId,
        title,
      })
      return normalizeSmallTask(smallTask)
    } catch (err) {
      throw errorProcessing(err, 'создании подзадачи')
    }
  }

  async updateSmallTask(id, { taskId, title }) {
    try {
      const smallTask = await apiClient.patch(API_ENDPOINTS.SMALL_TASKS.UPDATE(id), {
        task_id: taskId,
        title,
      })
      return normalizeSmallTask(smallTask)
    } catch (err) {
      throw errorProcessing(err, 'обновлении подзадачи')
    }
  }

  async deleteSmallTask(id) {
    try {
      return await apiClient.delete(API_ENDPOINTS.SMALL_TASKS.DELETE(id))
    } catch (err) {
      throw errorProcessing(err, 'удалении подзадачи')
    }
  }
}

export const smallTaskApi = new SmallTaskApi()
