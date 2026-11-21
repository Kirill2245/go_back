package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string      `json:"title" gorm:"not null"`
	Description string      `json:"description"`
	SmallTasks  []SmallTask `json:"small_tasks" gorm:"foreignKey:TaskID"`
}
