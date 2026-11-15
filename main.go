package main

import (
	task_controller "backToDo/internal/controller"
	"log"

	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()
	e.GET("/", task_controller.Hello)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
