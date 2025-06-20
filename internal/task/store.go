package task

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func SaveTasks(tasks *[]Task) error {
	jsonData, err := json.Marshal(*tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		log.Printf("failed to save task list\n")
		return err
	}
	log.Printf("succesfully saved task list\n")
	return nil
}

func AddTask(taskDesc string) error {
	var taskList *[]Task = LoadTasks()
	id := 0
	for _, v := range *taskList {
		if id <= v.Id {
			id = v.Id
		}
	}
	newTask := &Task{
		id + 1,
		taskDesc,
		todo,
		time.Now(),
		time.Now(),
	}
	*taskList = append(*taskList, *newTask)
	err := SaveTasks(taskList)
	log.Printf("added task \"%s\" with id %v \n", (*newTask).Description, (*newTask).Id)
	if err != nil {
		log.Printf("failed to add new task")
		return err
	}

	return nil
}
