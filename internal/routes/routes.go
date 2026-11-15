package routes

import (
	"net/http"

	"backToDo/internal/controller"

	"github.com/labstack/echo/v5"
)

func Register(e *echo.Echo, taskCtrl *controller.TaskController) {
	e.GET("/health", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Health - ok")
	})

	// Задачи
	e.POST("/create-task", taskCtrl.Create)
	e.GET("/get-tasks", taskCtrl.GetFullTasks)
	e.PATCH(("/update/:id"), taskCtrl.UpdateTask)
}
