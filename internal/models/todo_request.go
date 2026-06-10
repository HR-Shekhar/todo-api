package models

type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required,notblank"`
	Description string `json:"description"`
}