package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	todo       = "todo"
	inprogress = "in-progress"
	done       = "done"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func LoadTasks() *[]Task {
	var Tasks []Task

	file, err := os.OpenFile("tasks.json", os.O_RDONLY|os.O_CREATE, 0644)
	info, _ := os.Stat("tasks.json")
	if info.Size() == 0 {
		return &Tasks
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("failed to load task list")
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(data, &Tasks)
	return &Tasks
}

func UpdateTask(id int, newTascDesc string) error {
	taskList := LoadTasks()
	if len(*taskList) == 0 {
		return errors.New("there's no current tasks")
	}
	for i, _ := range *taskList {
		if id == (*taskList)[i].Id {
			(*taskList)[i].Description = newTascDesc
			(*taskList)[i].UpdatedAt = time.Now()
			err := SaveTasks(taskList)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("no task with such id")
}

func DeleteTask(id int) error {
	taskList := LoadTasks()
	if len(*taskList) == 1 {
		if (*taskList)[0].Id != id {
			return errors.New("no task with such id")
		}
		taskList = &[]Task{}
		err := SaveTasks(taskList)
		if err != nil {
			return err
		}
		return nil
	} else if len(*taskList) == 1 {
		return errors.New("there's no current tasks")
	}
	for i, v := range *taskList {
		if id == v.Id {
			*taskList = append((*taskList)[:i], (*taskList)[i+1:]...)
			err := SaveTasks(taskList)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("no task with such id")
}

func ListTasks(status string) *[]Task {
	tasks := LoadTasks()
	switch status {
	case "all":
		return tasks
	case "done":
		var doneTasks []Task
		for _, v := range *tasks {
			if v.Status == done {
				doneTasks = append(doneTasks, v)
			}
		}
		*tasks = doneTasks
	case "todo":
		fmt.Println("Heyyaa!")
		var todoTasks []Task
		for _, v := range *tasks {
			if v.Status == todo {
				todoTasks = append(todoTasks, v)
			}
		}
		*tasks = todoTasks
	case "in-progress":
		var inprgTasks []Task
		for _, v := range *tasks {
			if v.Status == inprogress {
				inprgTasks = append(inprgTasks, v)
			}
		}
	default:
		return nil
	}
	return tasks
}

func MarkInProgress(id int) (error, string) {
	tasks := LoadTasks()
	var taskIndex int
	for i, _ := range *tasks {
		if (*tasks)[i].Id == id {
			(*tasks)[i].Status = inprogress
			(*tasks)[i].UpdatedAt = time.Now()
			taskIndex = i
		}
	}
	err := SaveTasks(tasks)
	if err != nil {
		return err, fmt.Sprintf("couldn't save changes in task %v - %s", id, (*tasks)[taskIndex].Description)
	}
	return nil, fmt.Sprintf("Task %v has been marked in-progress!", id)
}

func MarkDone(id int) (error, string) {
	tasks := LoadTasks()
	var taskIndex int
	for i, _ := range *tasks {
		if (*tasks)[i].Id == id {
			(*tasks)[i].Status = done
			(*tasks)[i].UpdatedAt = time.Now()
			taskIndex = i
		}
	}
	err := SaveTasks(tasks)
	if err != nil {
		return err, fmt.Sprintf("couldn't save changes in task %v - %s", id, (*tasks)[taskIndex].Description)
	}
	return nil, fmt.Sprintf("Task %v has been marked done!", id)
}
