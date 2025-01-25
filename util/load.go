package util

import (
	"encoding/json"
	"log"
	"os"
	"task-tracker/model"
)

// Load function
// load tasks from json file
func Load() {
	file, err := os.Open("./tasks.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&model.Tasks)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Tasks loaded successfully")
}
