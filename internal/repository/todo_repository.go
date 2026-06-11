package repository

import (
	"context"
	"errors"

	"github.com/HR-Shekhar/todo-api/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoRepository struct {
	db *pgxpool.Pool
}

func (r *TodoRepository) CreateTodo(ctx context.Context, todo *models.Todo) (*models.Todo, error) {
	query := `INSERT INTO todos (
    user_id,
    title,
    description,
    completed
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
		todo.UserID,
		todo.Title,
		todo.Description,
		todo.Completed,
	)
	err := row.Scan(
		&todo.ID,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) GetTodo(
		ctx context.Context,
		todoID uuid.UUID,
		userID uuid.UUID,
	) (*models.Todo, error) {
	query := `
	SELECT
		id,
		user_id,
		title,
		description,
		completed,
		created_at,
		updated_at
	FROM todos
	WHERE id = $1
	AND user_id = $2;
	`
	
	row := r.db.QueryRow(
		ctx,
		query,
		todoID,
		userID,
	)

	todo := &models.Todo{}

	err := row.Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Title,
		&todo.Description,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTodoNotFound
		}
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) ListTodos(
	ctx context.Context,
	userID uuid.UUID,
	) ([]models.Todo, error) {
	query := `
	SELECT
		id,
		user_id,
		title,
		description,
		completed,
		created_at,
		updated_at
	FROM todos
	WHERE user_id = $1
	ORDER BY created_at DESC;
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(
			&todo.ID,
			&todo.UserID,
			&todo.Title,
			&todo.Description,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	// if an err occurs midway rows.Next() still gives false,
	// how to check that is it an err or no more rows, so we use rows.Err()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
} 

func (r *TodoRepository) UpdateTodo(
	ctx context.Context,
	todo *models.Todo,
	) (*models.Todo, error) {
	query := `
	UPDATE todos
	SET
		title = $1,
		description = $2,
		completed = $3,
		updated_at = NOW()
	WHERE id = $4
	AND user_id = $5
	RETURNING
		id,
		user_id,
		title,
		description,
		completed,
		created_at,
		updated_at;
	`
	row := r.db.QueryRow(
		ctx,
		query,
		todo.Title,
		todo.Description,
		todo.Completed,
		todo.ID,
		todo.UserID,
	)
	err := row.Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Title,
		&todo.Description,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTodoNotFound
		}
		return nil, err
	}
	return todo, nil 
}


func (r *TodoRepository) DeleteTodo(
		ctx context.Context,
		todoID uuid.UUID,
		userID uuid.UUID,
	) (error) {
	query := `
	DELETE FROM todos WHERE id = $1 AND user_id = $2
	`
	commandTag, err := r.db.Exec(ctx, query, todoID, userID)
	if err != nil {
		return err
	}
	msg := commandTag.RowsAffected()

	if msg == 0 {
		return ErrDeletionFailed
	}
	return nil
}