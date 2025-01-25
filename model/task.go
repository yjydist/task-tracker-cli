package model

// Task struct
type Task struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

var Tasks []Task
