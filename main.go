package main

import (
	controller "backToDo/internal/controller"

	"backToDo/internal/database"
	"backToDo/internal/repository"
	"backToDo/internal/routes"
	"backToDo/internal/service"
	"log"

	"github.com/labstack/echo/v5"
)

func main() {

	db := database.Connect()
	taskRepo := repository.NewTaskRepository(db)
	taskSvc := service.NewTaskService(taskRepo)
	taskController := controller.NewTaskController(taskSvc)

	e := echo.New()
	routes.Register(e, taskController)
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
