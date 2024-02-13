package main

import (
	"CLI_taskmanager/database"
	"CLI_taskmanager/task"
	"bufio"
	"fmt"

	// "fmt"
	"os"
	"os/signal"
)

func inputTasks() string {
	fmt.Println("input tasks: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	description := scanner.Text()

	return description
}

func main() {
	db := database.Database_connection()
	fmt.Println("connected to the databse")

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		defer sqlDB.Close()
		fmt.Println("Database connection closed")
	}()

	taskChan := make(chan string, 1)
	done := make(chan bool, 1)
	taskManager := task.NewMananger(taskChan, done)
	taskManager.DisplayList()
	go func() {
		var continuee int
		continuee = 1
		for continuee == 1 {
			fmt.Println("Enter 1 for input of task and 2 for the deletion of task")
			var choose int
			fmt.Scanf("%d", &choose)
			switch choose {
			case 1:
				taskChan <- inputTasks()
				break

			case 2:
				taskManager.DeleteTask()
				break
			default:
				fmt.Println("wrong choice please try it again")
			}
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		done <- true
	}()
	taskManager.ListernForTasks()
}
