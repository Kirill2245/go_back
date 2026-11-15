package task_controller

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func Hello(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, Echo v5")
}
