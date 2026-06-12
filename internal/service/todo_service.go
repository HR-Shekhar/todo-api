package service

import (
	"context"
	"errors"

	"github.com/HR-Shekhar/todo-api/internal/models"
	"github.com/HR-Shekhar/todo-api/internal/repository"
	"github.com/google/uuid"
)

type TodoService struct {
	todoRepo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{
		todoRepo: repo,
	}
}

func (s *TodoService) CreateTodo(
	ctx context.Context,
	req *models.CreateTodoRequest,
	userID uuid.UUID) (*models.Todo, error) {
	todo := &models.Todo{
		UserID: userID,
		Title: req.Title,
		Description: req.Description,
	}
	created, err := s.todoRepo.CreateTodo(ctx, todo)

	if err != nil {
		return nil, ErrTodoCreationFailed		
	}

	return created, nil
}

func (s *TodoService) UpdateTodo(
	ctx context.Context,
	req *models.UpdateTodoRequest,
	todoID uuid.UUID,
	userID uuid.UUID,
	) (*models.Todo, error) {
	todo := &models.Todo{
		ID: todoID,
		UserID: userID,
		Title: req.Title,
		Description: req.Description,
		Completed: req.Completed,
	}
	updated, err := s.todoRepo.UpdateTodo(ctx, todo)

	if err != nil {
		return nil, ErrTodoNotFound
	}

	return updated, nil
}

func (s *TodoService) ListTodos(
	ctx context.Context,
	userID uuid.UUID,
	) ([]models.Todo, error) {
	todos, err := s.todoRepo.ListTodos(ctx, userID)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) GetTodo(
	ctx context.Context,
	todoID uuid.UUID,
	userID uuid.UUID,
	) (*models.Todo, error) {
	todo, err := s.todoRepo.GetTodo(ctx, todoID, userID)

	if err != nil {
		if errors.Is(err, repository.ErrTodoNotFound) {
			return nil, ErrTodoNotFound
		}
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) DeleteTodo(
	ctx context.Context,
	todoID uuid.UUID,
	userID uuid.UUID,
	) (error) {
	err := s.todoRepo.DeleteTodo(ctx, todoID, userID)

	if err != nil {
		if errors.Is(err, repository.ErrTodoNotFound) {
			return ErrTodoNotFound
		}
		return err
	}

	return nil
}