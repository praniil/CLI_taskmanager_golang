package task

import (
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string
}

func NewTask(description string) *Task {
	return &Task{
		Description: description,
	}
}

func (t *Task) String() string {
	return t.Description
}
