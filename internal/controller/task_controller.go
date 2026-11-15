package controller

import (
	"net/http"
	"strconv"

	"backToDo/internal/service"

	"github.com/labstack/echo/v5"
)

type TaskController struct {
	svc *service.TaskService
}

type TaskRequest struct {
	Title       string `json:"title" example:"Buy milk"`
	Description string `json:"description" example:"2 liters of milk"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"invalid body"`
}

func NewTaskController(svc *service.TaskService) *TaskController {
	return &TaskController{svc: svc}
}

// Create creates a new task.
//
//	@Summary		Create a task
//	@Description	Creates a task with the given title and description
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			request	body		controller.TaskRequest	true	"Task payload"
//	@Success		201		{object}	model.Task
//	@Failure		400		{object}	controller.ErrorResponse
//	@Failure		500		{object}	controller.ErrorResponse
//	@Router			/create-task [post]
func (h *TaskController) Create(c *echo.Context) error {
	var req TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid body"})
	}

	task, err := h.svc.CreateTask(c.Request().Context(), req.Title, req.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusCreated, task)
}

// GetFullTasks returns all tasks.
//
//	@Summary		List tasks
//	@Description	Returns all tasks
//	@Tags			tasks
//	@Produce		json
//	@Success		200	{array}		model.Task
//	@Failure		500	{object}	controller.ErrorResponse
//	@Router			/get-tasks [get]
func (h *TaskController) GetFullTasks(c *echo.Context) error {
	tasks, err := h.svc.GetFullTasks(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, tasks)
}

// UpdateTask updates an existing task.
//
//	@Summary		Update a task
//	@Description	Updates the title and description of a task by ID
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"Task ID"
//	@Param			request	body		controller.TaskRequest	true	"Task payload"
//	@Success		200		{object}	model.Task
//	@Failure		400		{object}	controller.ErrorResponse
//	@Failure		404		{object}	controller.ErrorResponse
//	@Failure		500		{object}	controller.ErrorResponse
//	@Router			/update/{id} [patch]
func (h *TaskController) UpdateTask(c *echo.Context) error {
	var req TaskRequest

	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid id"})
	}
	id := uint(id64)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid body"})
	}

	task, err := h.svc.UpdateTask(c.Request().Context(), id, req.Title, req.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, task)
}
