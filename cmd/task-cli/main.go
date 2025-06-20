package main

import (
	task "cli-task-tracker/internal/task"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	var action string
	if len(os.Args) >= 2 {
		action = os.Args[1]
	}
	switch action {
	case "add":
		err := task.AddTask(os.Args[2])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "update":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Task id %s is incorrect type", os.Args[2])
			os.Exit(1)
		}
		err = task.UpdateTask(id, os.Args[3])
	case "delete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("incorrect value type")
		}
		err = task.DeleteTask(id)
		if err != nil {
			fmt.Println(err)
		}
	case "mark-in-progress":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err, result := task.MarkInProgress(id)
		fmt.Println(result)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "mark-done":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err, result := task.MarkDone(id)
		fmt.Println(result)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "list":
		if len(os.Args) > 3 {
			fmt.Println("incorrect number of variables")
			os.Exit(1)
		}
		if len(os.Args) == 2 {
			tasks := task.ListTasks("all")
			fmt.Println(tasks)
		} else {
			tasks := task.ListTasks(os.Args[2])
			fmt.Println(tasks)
		}
	}
	os.Exit(1)
}
