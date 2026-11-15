package controller

import (
	"net/http"

	"backToDo/internal/service"

	"strconv"

	"github.com/labstack/echo/v5"
)

type TaskController struct {
	svc *service.TaskService
}

func NewTaskController(svc *service.TaskService) *TaskController {
	return &TaskController{svc: svc}
}

func (h *TaskController) Create(c *echo.Context) error {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body"})
	}

	task, err := h.svc.CreateTask(c.Request().Context(), req.Title, req.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// 3. Отвечаем
	return c.JSON(http.StatusCreated, task)
}

func (h *TaskController) GetFullTasks(c *echo.Context) error {
	tasks, err := h.svc.GetFullTasks(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskController) UpdateTask(c *echo.Context) error {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	id := uint(id64)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body"})
	}

	task, err := h.svc.UpdateTask(c.Request().Context(), id, req.Title, req.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, task)
}
