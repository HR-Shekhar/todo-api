package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID uuid.UUID	    	`json:"id"`
	UserID uuid.UUID	    	`json:"user_id"`
	Title string 		`json:"title"`
	Description string  `json:"description"`
	IsCompleted bool    `json:"is_completed"`
	CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
