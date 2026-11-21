export const API_ENDPOINTS = {
  TASKS: {
    LIST: '/get-tasks',
    CREATE: '/create-task',
    UPDATE: (id) => `/update/${id}`,
    DELETE: (id) => `/delete-task/${id}`,
  },
  SMALL_TASKS: {
    LIST: (taskID) => `/get-small-tasks/${taskID}`,
    CREATE: '/create-small-task',
    UPDATE: (id) => `/update-small-task/${id}`,
    DELETE: (id) => `/delete-small-task/${id}`,
  },
}
