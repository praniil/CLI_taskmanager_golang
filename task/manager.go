package task

import (
	"CLI_taskmanager/database"
	// "gorm.io/driver/postgres"
	"bufio"
	"fmt"
	"os"
)

type ManagerStruct struct {
	tasks    []*Task
	taskChan chan string
	done     chan bool
}

func NewMananger(taskChan chan string, done chan bool) *ManagerStruct {
	return &ManagerStruct{
		tasks:    []*Task{},
		taskChan: taskChan,
		done:     done,
	}
}

func (m *ManagerStruct) AddTask(description string) {
	db := database.Database_connection()
	db.AutoMigrate(&Task{})
	m.tasks = append(m.tasks, NewTask(description))
	task := &Task{
		Description: description,
	}
	result := db.Create(task)
	if result.Error != nil {
		fmt.Println("failed to create a task or add task in the database")
	}
	var taskInfo []*Task
	reslt := db.Find(&taskInfo)
	if reslt.Error != nil {
		fmt.Println("couldnot get the record from the database", reslt.Error)
	}
	fmt.Println(taskInfo)
}

func (m *ManagerStruct) DeleteTask() {
	fmt.Println("Enter the id you want to delete: ")
	var delete int
	fmt.Scanf("%d", &delete)
	db := database.Database_connection()
	result := db.Delete(&Task{}, delete)
	if result.Error != nil {
		fmt.Println("failed to delete the record: ", result.Error)
	}
	rowsDeleted := result.RowsAffected
	fmt.Printf("no of rows deleted: %d \n", rowsDeleted)
	var taskInfo []*Task
	reslt := db.Find(&taskInfo)
	if reslt.Error != nil {
		fmt.Println("couldnot get the record from the database", reslt.Error)
	}
	fmt.Println(taskInfo)
}

func UpdateTask() {
	db := database.Database_connection()
	fmt.Printf("Enter the id you want to update: ")
	var id int
	fmt.Scanf("%d", &id)
	fmt.Println("input tasks: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	description := scanner.Text()
	result := db.Model(&Task{}).Where("id =?", id).Update("Description", description)
	rowsUpdated := result.RowsAffected
	fmt.Printf("no of rows updated: %d\n", rowsUpdated)

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
func (m *ManagerStruct) DisplayList() []*Task {
	return m.tasks
}
