package routes

import (
	"net/http"

	"backToDo/internal/controller"

	"github.com/labstack/echo/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Register(e *echo.Echo, taskCtrl *controller.TaskController, smallTaskCtrl *controller.SmallTaskController) {
	e.GET("/health", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Health - ok")
	})

	e.GET("/swagger", func(c *echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	e.GET("/swagger/*", echo.WrapHandler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)))

	e.POST("/create-task", taskCtrl.Create)
	e.GET("/get-tasks", taskCtrl.GetFullTasks)
	e.PATCH("/update/:id", taskCtrl.UpdateTask)
	e.DELETE("/delete-task/:id", taskCtrl.DeleteTask)

	e.POST("/create-small-task", smallTaskCtrl.Create)
	e.GET("/get-small-tasks/:taskID", smallTaskCtrl.GetSmallTasks)
	e.PATCH("/update-small-task/:id", smallTaskCtrl.UpdateSmallTask)
	e.DELETE("/delete-small-task/:id", smallTaskCtrl.DeleteSmallTask)
}
