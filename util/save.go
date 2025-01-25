package util

import (
	"encoding/json"
	"log"
	"os"
	"task-tracker/model"
)

// Save function
// save tasks to json file
func Save() {
	file, err := os.Open("./tasks.json")
	if err != nil {
		return
	}
	tasksJson, err := json.Marshal(model.Tasks)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = file.Write(tasksJson)
	if err != nil {
		log.Fatal(err)
		return
	}
}
