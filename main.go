//go:generate swag init -g main.go -o docs --parseDependency --parseInternal

package main

import (
	"log"
	"os"

	_ "backToDo/docs"
	"backToDo/internal/controller"
	"backToDo/internal/database"
	"backToDo/internal/repository"
	"backToDo/internal/routes"
	"backToDo/internal/service"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// @title			BackToDo API
// @version		1.0
// @description	REST API for managing tasks
// @host			localhost:8080
// @BasePath		/
func main() {
	db := database.Connect()
	taskRepo := repository.NewTaskRepository(db)
	taskSvc := service.NewTaskService(taskRepo)
	taskController := controller.NewTaskController(taskSvc)

	e := echo.New()

	e.Use(middleware.CORS(os.Getenv("ALLOWORIGIN")))
	routes.Register(e, taskController)
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
