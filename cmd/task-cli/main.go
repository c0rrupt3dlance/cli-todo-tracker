package main

import (
	taskerror "cli-task-tracker/internal/errors"
	"cli-task-tracker/internal/task"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var action string
	if len(os.Args) >= 2 {
		action = os.Args[1]
	}
	switch action {
	case "add":
		err := taskerror.CheckAddArgs(os.Args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = task.AddTask(os.Args[2])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "update":
		err := taskerror.CheckUpdateArgs(os.Args)
		if err != nil {
			fmt.Println(err)
		}
		id, _ := strconv.Atoi(os.Args[2])
		err = task.UpdateTask(id, os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "delete":
		err := taskerror.CheckDeleteArgs(os.Args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		id, _ := strconv.Atoi(os.Args[2])
		err = task.DeleteTask(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "mark-in-progress":
		err := taskerror.CheckStatusArgs(os.Args, "mark-in-progress")
		id, _ := strconv.Atoi(os.Args[2])
		_, result := task.MarkInProgress(id)
		fmt.Println(result)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "mark-done":
		err := taskerror.CheckStatusArgs(os.Args, "mark-in-progress")
		id, _ := strconv.Atoi(os.Args[2])
		_, result := task.MarkDone(id)
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
			for _, v := range *tasks {
				fmt.Println("ID:", v.Id)
				fmt.Println("Description:", v.Description)
				fmt.Println("Status:", v.Status)
				fmt.Println("Created at:", v.CreatedAt)
				fmt.Println("Updated at:", v.UpdatedAt)
				fmt.Println("========================================")
			}
		} else {
			tasks := task.ListTasks(os.Args[2])
			for _, v := range *tasks {
				fmt.Println("ID:", v.Id)
				fmt.Println("Description:", v.Description)
				fmt.Println("Status:", v.Status)
				fmt.Println("Created at:", v.CreatedAt)
				fmt.Println("Updated at:", v.UpdatedAt)
			}
		}
	}
	os.Exit(1)
}
