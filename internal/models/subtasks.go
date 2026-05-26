package models

import (
	"time"
)

type Subtask struct {
	ID int			    `json:"id"`
	TodoID string		`json:"todo_id"`
	Title string 		`json:"title"`
	IsCompleted bool    `json:"is_completed"`
	CreatedAt time.Time `json:"created_at"`
}
