package main

import (
	"CLI_taskmanager/database"
	tm "CLI_taskmanager/task"
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
	taskManager := tm.NewMananger(taskChan, done)
	taskManager.DisplayList()
	go func() {
		var continuee int
		continuee = 1
		for continuee == 1 {
			fmt.Printf("Select\n 1. for input of task\n 2. for the deletion of task\n 3. for updating the task\n")
			var choose int
			fmt.Scanf("%d", &choose)
			switch choose {
			case 1:
				taskChan <- inputTasks()
				break

			case 2:
				taskManager.DeleteTask()
				break

			case 3:
				tm.UpdateTask()
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
