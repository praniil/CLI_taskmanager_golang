package task

import (
	"CLI_taskmanager/database"
	"fmt"
)

type ManagerStruct struct {
	tasks    []*TaskStruct
	taskChan chan string
	done     chan bool
}

func NewMananger(taskChan chan string, done chan bool) *ManagerStruct {
	return &ManagerStruct{
		tasks:    []*TaskStruct{},
		taskChan: taskChan,
		done:     done,
	}
}

func (m *ManagerStruct) AddTask(description string) {
	db := database.Database_connection()
	m.tasks = append(m.tasks, NewTask(description))
	db.AutoMigrate(&TaskStruct{})
	task := *&TaskStruct{
		Description: description,
	}
	result := db.Create(task)
	if result.Error != nil {
		fmt.Println("failed to create a task or add task in the database")
	}
	fmt.Println(m.tasks)
}

func (m *ManagerStruct) ListernForTasks() {
	for {
		select {
		case description := <-m.taskChan:
			m.AddTask(description)
		case <-m.done:
			return
		}
	}
}
func (m *ManagerStruct) DisplayList() []*TaskStruct {
	return m.tasks
}
