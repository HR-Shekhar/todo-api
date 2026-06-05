package service

import (
	"context"
	"strings"
	"errors"
	
    "github.com/jackc/pgx/v5/pgconn"

	"github.com/HR-Shekhar/todo-api/internal/models"
	"github.com/HR-Shekhar/todo-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
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
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	// build user from request
	user := &models.User{
		FullName: req.FullName,
		Username: strings.TrimSpace(strings.ToLower(req.Username)),
		Email:    strings.TrimSpace(strings.ToLower(req.Email)),
		PasswordHash: string(hash),
	}

	created, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
				case "users_email_key":
					return nil, ErrEmailAlreadyExists
				case "users_username_key":
					return nil, ErrUsernameAlreadyExists
			}
		}
		return nil, err
	}

	return created, nil
}