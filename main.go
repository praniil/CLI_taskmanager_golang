package main

import (
	taskmanager "CLI_taskmanager/task"
	"bufio"
	"fmt"

	// "fmt"
	"os"
	"os/signal"
)

func main() {
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
