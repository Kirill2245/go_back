export const parseServerError = (error) => {
  const result = { fields: {} }

  if (error && typeof error === 'object' && 'isAxiosError' in error) {
    const axiosError = error

    if (axiosError.response) {
      const statusCode = axiosError.response.status
      const serverData = axiosError.response.data
      result.statusCode = statusCode

      if (serverData?.message) {
        result.general = Array.isArray(serverData.message)
          ? serverData.message.join('\n')
          : serverData.message
      } else if (serverData?.error) {
        result.general = serverData.error
      } else {
        result.general = `Ошибка сервера: ${statusCode}`
      }
    } else if (axiosError.request) {
      result.general = 'Нет соединения с сервером. Проверьте интернет-соединение.'
    } else {
      result.general = axiosError.message || 'Произошла неизвестная ошибка'
    }
  } else if (error instanceof Error) {
    result.general = error.message
  } else {
    result.general = 'Произошла неизвестная ошибка'
  }

  return result
}

export const errorProcessing = (error, label) => {
  const parsed = parseServerError(error)
  const enhancedError = new Error(parsed.general || `Произошла ошибка при ${label}`)

  enhancedError.originalError = error
  enhancedError.statusCode = parsed.statusCode

  if (error && typeof error === 'object' && 'response' in error) {
    enhancedError.serverData = error.response?.data
  }

  return enhancedError
}
