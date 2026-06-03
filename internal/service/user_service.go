package service

import (
	"context"
	"os/user"

	"github.com/HR-Shekhar/todo-api/internal/models"
	"github.com/HR-Shekhar/todo-api/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (s *UserService) RegisterUser(
    ctx context.Context,
    req *models.CreateUserRequest,
) (*models.User, error) {
	s.userRepo.CreateUser(ctx, &models.User{})
}