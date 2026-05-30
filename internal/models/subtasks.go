package models

import (
	"time"

	"github.com/google/uuid"
)

type Subtask struct {
	ID uuid.UUID 	    `json:"id"`
	TodoID uuid.UUID	`json:"todo_id"`
	Title string 		`json:"title"`
	IsCompleted bool    `json:"is_completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
