package repository

import (
	"backToDo/internal/model"
	"context"

	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(ctx context.Context, task *model.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *taskRepository) GetByID(ctx context.Context, id uint) (*model.Task, error) {
	var task model.Task
	err := r.db.WithContext(ctx).First(&task, id).Error
	return &task, err
}

func (r *taskRepository) FindTasks(ctx context.Context) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.WithContext(ctx).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *model.Task) error {
	return r.db.WithContext(ctx).Model(task).Updates(model.Task{
		Title:       task.Title,
		Description: task.Description,
	}).Error
}
