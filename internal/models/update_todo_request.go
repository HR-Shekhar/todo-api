package models

type UpdateTodoRequest struct {
	Title       string `json:"title" validate:"required,notblank"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}