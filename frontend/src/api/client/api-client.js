import axios from 'axios'

import { API_CONFIG } from './api-config'
import { parseServerError } from '../helpers/error-processing'

function createApiError(message, init) {
  const err = new Error(message)
  err.name = 'ApiError'
  Object.defineProperty(err, 'parsed', { value: init.parsed, enumerable: true })
  Object.defineProperty(err, 'originalError', {
    value: init.originalError,
    enumerable: true,
  })

  if (init.statusCode !== undefined) {
    Object.defineProperty(err, 'statusCode', {
      value: init.statusCode,
      enumerable: true,
    })
  }

  if (init.code !== undefined) {
    Object.defineProperty(err, 'code', { value: init.code, enumerable: true })
  }

  return err
}

export function isApiError(error) {
  if (!(error instanceof Error) || error.name !== 'ApiError') return false

  return (
    error.parsed !== undefined &&
    typeof error.parsed === 'object' &&
    error.parsed !== null &&
    'originalError' in error
  )
}

class ApiClient {
  constructor() {
    this.instance = axios.create({
      baseURL: API_CONFIG.baseURL,
      timeout: API_CONFIG.timeout,
      headers: API_CONFIG.headers,
    })

    this.setupInterceptors()
  }

  buildRequestUrl(config) {
    try {
      return this.instance.getUri(config)
    } catch {
      const base = (config.baseURL ?? this.instance.defaults.baseURL ?? '').replace(/\/$/, '')
      const path = config.url ?? ''
      if (!base) return path || ''
      if (!path) return base
      return path.startsWith('http') ? path : `${base}/${path.replace(/^\//, '')}`
    }
  }

  normalizeFailure(error) {
    if (isApiError(error)) {
      return error
    }

    if (!axios.isAxiosError(error)) {
      const parsed = parseServerError(error)
      const message = parsed.general || 'Произошла неизвестная ошибка'
      return createApiError(message, {
        parsed,
        originalError: error,
        statusCode: parsed.statusCode,
      })
    }

    const axiosError = error
    let parsed = parseServerError(axiosError)

    if (axiosError.code === 'ECONNABORTED') {
      parsed = {
        ...parsed,
        general: 'Превышено время ожидания ответа сервера',
      }
    }

    if (axiosError.code === 'ERR_NETWORK') {
      parsed = {
        ...parsed,
        general:
          parsed.general ||
          `Сервер недоступен. Проверьте подключение и что backend запущен на ${API_CONFIG.baseURL}`,
      }
    }

    const message = parsed.general || axiosError.message || 'Произошла неизвестная ошибка'

    return createApiError(message, {
      parsed,
      originalError: axiosError,
      statusCode: axiosError.response?.status ?? parsed.statusCode,
      code: axiosError.code,
    })
  }

  setupInterceptors() {
    this.instance.interceptors.request.use(
      (config) => {
        if (import.meta.env.DEV) {
          const uri = this.buildRequestUrl(config)
          console.log(`📡 ${config.method?.toUpperCase() ?? 'GET'} ${uri}`)
        }

        return config
      },
      (error) => Promise.reject(this.normalizeFailure(error)),
    )

    this.instance.interceptors.response.use(
      (response) => response,
      (error) => Promise.reject(this.normalizeFailure(error)),
    )
  }

  async get(url, config) {
    const response = await this.instance.get(url, config)
    return response.data
  }

  async post(url, data, config) {
    const response = await this.instance.post(url, data, config)
    return response.data
  }

  async patch(url, data, config) {
    const response = await this.instance.patch(url, data, config)
    return response.data
  }

  async delete(url, config) {
    const response = await this.instance.delete(url, config)
    return response.data
  }
}

export const apiClient = new ApiClient()
