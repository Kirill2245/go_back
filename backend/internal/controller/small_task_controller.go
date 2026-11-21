package controller

import (
	"errors"
	"net/http"
	"strconv"

	"backToDo/internal/service"

	"github.com/labstack/echo/v5"
)

type SmallTaskController struct {
	svc *service.TaskService
}

type SmallTaskRequest struct {
	TaskID uint   `json:"task_id" example:"1"`
	Title  string `json:"title" example:"Buy milk"`
}

func NewSmallTaskController(svc *service.TaskService) *SmallTaskController {
	return &SmallTaskController{svc: svc}
}

// Create creates a new small task for the given task.
//
//	@Summary		Create a small task
//	@Description	Creates a small task with the given title for a task
//	@Tags			small-tasks
//	@Accept			json
//	@Produce		json
//	@Param			request	body		controller.SmallTaskRequest	true	"Small task payload"
//	@Success		201		{object}	model.SmallTask
//	@Failure		400		{object}	controller.ErrorResponse
//	@Failure		404		{object}	controller.ErrorResponse
//	@Failure		500		{object}	controller.ErrorResponse
//	@Router			/create-small-task [post]
func (h *SmallTaskController) Create(c *echo.Context) error {
	var req SmallTaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid body"})
	}

	smallTask, err := h.svc.CreateSmallTask(c.Request().Context(), req.TaskID, req.Title)
	if err != nil {
		if errors.Is(err, service.ErrTitleRequired) {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		}
		if errors.Is(err, service.ErrTaskNotFound) {
			return c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusCreated, smallTask)
}

// GetSmallTasks returns all small tasks of a task.
//
//	@Summary		List small tasks
//	@Description	Returns all small tasks of a task by task ID
//	@Tags			small-tasks
//	@Produce		json
//	@Param			taskID	path		int	true	"Task ID"
//	@Success		200		{array}		model.SmallTask
//	@Failure		400		{object}	controller.ErrorResponse
//	@Failure		404		{object}	controller.ErrorResponse
//	@Failure		500		{object}	controller.ErrorResponse
//	@Router			/get-small-tasks/{taskID} [get]
func (h *SmallTaskController) GetSmallTasks(c *echo.Context) error {
	taskID64, err := strconv.ParseUint(c.Param("taskID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid task id"})
	}

	smallTasks, err := h.svc.GetSmallTasks(c.Request().Context(), uint(taskID64))
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			return c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, smallTasks)
}

// UpdateSmallTask updates an existing small task.
//
//	@Summary		Update a small task
//	@Description	Updates the title of a small task by ID
//	@Tags			small-tasks
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Small task ID"
//	@Param			request	body		controller.SmallTaskRequest	true	"Small task payload"
//	@Success		200		{object}	model.SmallTask
//	@Failure		400		{object}	controller.ErrorResponse
//	@Failure		404		{object}	controller.ErrorResponse
//	@Failure		500		{object}	controller.ErrorResponse
//	@Router			/update-small-task/{id} [patch]
func (h *SmallTaskController) UpdateSmallTask(c *echo.Context) error {
	var req SmallTaskRequest

	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid id"})
	}
	id := uint(id64)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid body"})
	}

	smallTask, err := h.svc.UpdateSmallTask(c.Request().Context(), id, req.Title)
	if err != nil {
		if errors.Is(err, service.ErrTitleRequired) {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		}
		if errors.Is(err, service.ErrSmallTaskNotFound) {
			return c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, smallTask)
}

// DeleteSmallTask deletes a small task by ID.
//
//	@Summary		Delete a small task
//	@Description	Deletes a small task by ID
//	@Tags			small-tasks
//	@Produce		json
//	@Param			id	path		int	true	"Small task ID"
//	@Success		200	{object}	controller.MessageResponse
//	@Failure		400	{object}	controller.ErrorResponse
//	@Failure		404	{object}	controller.ErrorResponse
//	@Failure		500	{object}	controller.ErrorResponse
//	@Router			/delete-small-task/{id} [delete]
func (h *SmallTaskController) DeleteSmallTask(c *echo.Context) error {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid id"})
	}

	if err := h.svc.DeleteSmallTask(c.Request().Context(), uint(id64)); err != nil {
		if errors.Is(err, service.ErrSmallTaskNotFound) {
			return c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, MessageResponse{Message: "small task deleted"})
}
