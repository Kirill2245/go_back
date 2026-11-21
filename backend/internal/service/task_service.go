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
func (s *TaskService) taskExists(ctx context.Context, id uint) (*model.Task, error) {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return task, nil
}

func (s *TaskService) smallTaskExists(ctx context.Context, id uint) (*model.SmallTask, error) {
	smallTask, err := s.repo.GetSmallTaskByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return smallTask, nil
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

	task, err := s.taskExists(ctx, id)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, ErrTaskNotFound
	}
	task.Title = title
	task.Description = description

	if err := s.repo.UpdateTask(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id uint) error {
	exists, err := s.taskExists(ctx, id)
	if err != nil {
		return err
	}
	if exists == nil {
		return ErrTaskNotFound
	}
	return s.repo.RemoveTask(ctx, id)
}

func (s *TaskService) CreateSmallTask(ctx context.Context, taskID uint, title string) (*model.SmallTask, error) {
	if title == "" {
		return nil, ErrTitleRequired
	}

	task, err := s.taskExists(ctx, taskID)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, ErrTaskNotFound
	}

	smallTask := &model.SmallTask{Title: title, TaskID: taskID}
	if err := s.repo.CreateSmallTask(ctx, smallTask); err != nil {
		return nil, err
	}
	return smallTask, nil
}

func (s *TaskService) GetSmallTasks(ctx context.Context, taskID uint) ([]model.SmallTask, error) {
	task, err := s.taskExists(ctx, taskID)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, ErrTaskNotFound
	}
	return s.repo.FindSmallTasks(ctx, taskID)
}

func (s *TaskService) UpdateSmallTask(ctx context.Context, id uint, title string) (*model.SmallTask, error) {
	if title == "" {
		return nil, ErrTitleRequired
	}

	smallTask, err := s.smallTaskExists(ctx, id)
	if err != nil {
		return nil, err
	}
	if smallTask == nil {
		return nil, ErrSmallTaskNotFound
	}
	smallTask.Title = title

	if err := s.repo.UpdateSmallTask(ctx, smallTask); err != nil {
		return nil, err
	}

	return smallTask, nil
}

func (s *TaskService) DeleteSmallTask(ctx context.Context, id uint) error {
	exists, err := s.smallTaskExists(ctx, id)
	if err != nil {
		return err
	}
	if exists == nil {
		return ErrSmallTaskNotFound
	}
	return s.repo.RemoveSmallTask(ctx, id)
}
