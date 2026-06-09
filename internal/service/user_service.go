package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/HR-Shekhar/todo-api/internal/models"
	"github.com/HR-Shekhar/todo-api/internal/repository"
	
)

type UserService struct {
	userRepo *repository.UserRepository
	jwtSecret string
}

func NewUserService(repo *repository.UserRepository, jwtSecret string) *UserService {
	return &UserService{
		userRepo: repo,
		jwtSecret: jwtSecret,
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

func (s *UserService) LoginUser(
	ctx context.Context,
	req *models.LoginRequest,
) (string, error) {
	user, err := s.userRepo.GetUserByEmail(
		ctx,
		strings.TrimSpace(strings.ToLower(req.Email)),
	)

	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return "", ErrInvalidCredentials
		}
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash),[]byte(req.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", ErrInvalidCredentials
		}
		return "", err
	}
	token, err := s.generateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) generateToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID.String(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(
		[]byte(s.jwtSecret),
	)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}