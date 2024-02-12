package main

import (
	"CLI_taskmanager/database"
	taskmanager "CLI_taskmanager/task"
	"bufio"
	"fmt"

	// "fmt"
	"os"
	"os/signal"
)

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
	taskManager := taskmanager.NewMananger(taskChan, done)
	taskManager.DisplayList()
	go func() {
		for {
			fmt.Println("input tasks: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			description := scanner.Text()
			taskChan <- description
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
