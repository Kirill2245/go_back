package service

import "errors"

var (
	ErrTitleRequired = errors.New("title is required")
	ErrTaskNotFound  = errors.New("task not found")
)
