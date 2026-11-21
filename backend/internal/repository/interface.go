package repository

import (
	"backToDo/internal/model"
	"context"
)

type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) error
	GetByID(ctx context.Context, id uint) (*model.Task, error)
	FindTasks(ctx context.Context) ([]model.Task, error)
	UpdateTask(ctx context.Context, task *model.Task) error
	RemoveTask(ctx context.Context, id uint) error

	CreateSmallTask(ctx context.Context, smallTask *model.SmallTask) error
	GetSmallTaskByID(ctx context.Context, id uint) (*model.SmallTask, error)
	FindSmallTasks(ctx context.Context, taskID uint) ([]model.SmallTask, error)
	UpdateSmallTask(ctx context.Context, smallTask *model.SmallTask) error
	RemoveSmallTask(ctx context.Context, id uint) error
}
