package service

import (
	"backToDo/internal/model"
	"backToDo/internal/repository"

	"context"
	"errors"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, title, description string) (*model.Task, error) {

	if title == "" {
		return nil, echo.NewHTTPError(400, "name is required")
	}

	task := &model.Task{Title: title, Description: description}
	if err := s.repo.Create(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) GetFullTasks(ctx context.Context) ([]model.Task, error) {
	tasks, err := s.repo.FindTasks(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id uint, title, description string) (*model.Task, error) {
	if title == "" {
		return nil, ErrTitleRequired
	}

	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}

	task.Title = title
	task.Description = description

	if err := s.repo.UpdateTask(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}
