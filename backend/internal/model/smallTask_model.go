package model

import "gorm.io/gorm"

type SmallTask struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	TaskID uint   `json:"task_id"`
}
