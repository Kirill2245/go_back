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
	err := r.db.WithContext(ctx).Preload("SmallTasks").First(&task, id).Error
	return &task, err
}

func (r *taskRepository) FindTasks(ctx context.Context) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.WithContext(ctx).Preload("SmallTasks").Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *model.Task) error {
	return r.db.WithContext(ctx).Model(task).Updates(model.Task{
		Title:       task.Title,
		Description: task.Description,
	}).Error
}

func (r *taskRepository) RemoveTask(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("task_id = ?", id).Delete(&model.SmallTask{}).Error; err != nil {
			return err
		}
		return tx.Unscoped().Delete(&model.Task{}, id).Error
	})
}

func (r *taskRepository) CreateSmallTask(ctx context.Context, smallTask *model.SmallTask) error {
	return r.db.WithContext(ctx).Create(smallTask).Error
}

func (r *taskRepository) GetSmallTaskByID(ctx context.Context, id uint) (*model.SmallTask, error) {
	var smallTask model.SmallTask
	err := r.db.WithContext(ctx).First(&smallTask, id).Error
	return &smallTask, err
}

func (r *taskRepository) FindSmallTasks(ctx context.Context, taskID uint) ([]model.SmallTask, error) {
	var smallTasks []model.SmallTask
	err := r.db.WithContext(ctx).Where("task_id = ?", taskID).Find(&smallTasks).Error
	return smallTasks, err
}

func (r *taskRepository) UpdateSmallTask(ctx context.Context, smallTask *model.SmallTask) error {
	return r.db.WithContext(ctx).Model(smallTask).Updates(model.SmallTask{
		Title: smallTask.Title,
	}).Error
}

func (r *taskRepository) RemoveSmallTask(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Unscoped().Delete(&model.SmallTask{}, id).Error
}
