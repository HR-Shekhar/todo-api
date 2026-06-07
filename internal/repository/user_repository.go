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

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := `INSERT INTO users (
    full_name,
    username,
    email,
    password_hash
	)
	VALUES (
		$1,
		$2,
		$3,
		$4
	)
	RETURNING
		id,
		created_at,
		updated_at;`
		
	row := r.db.QueryRow(
		ctx,
		query,
		user.FullName,
		user.Username,
		user.Email,
		user.PasswordHash,
	)

	err := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	
	if err != nil {
		return nil,err
	}

	return user,nil
}

func (r *UserRepository) GetUserByEmail(
    ctx context.Context,
    email string,
) (*models.User, error) {

}