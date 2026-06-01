package repository

import (
	"context"
	"github.com/HR-Shekhar/todo-api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,	
	}
}

func (r *UserRepository) CreateUser(user *models.User, ctx context.Context) (*models.User, error) {

}