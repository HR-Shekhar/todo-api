package handler

import (
	"github.com/HR-Shekhar/todo-api/internal/repository"
)

type UserHandler struct {
	userRepo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: repo,
	}
}


func (h *UserHandler) CreateUser(...)