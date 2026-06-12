package handler

import (
	"net/http"

	"github.com/HR-Shekhar/todo-api/internal/models"
	"github.com/HR-Shekhar/todo-api/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	todoService *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: service,
	}
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
	var req models.CreateTodoRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	if err := c.Validate(&req); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)

		if !ok {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		var messages []string

		for _, validationErr := range validationErrors {
			switch validationErr.Tag() {
				case "required":
					messages = append(messages, validationErr.Field()+" is required.")
				case "notblank":
					messages = append(messages, validationErr.Field()+" should not be blank.")
			}
		}
		return c.JSON(
			http.StatusBadRequest,
			map[string]any{
				"errors": messages,
			},
		)
	}
	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": "failed to get authenticated user",
			},
		)
	}

	parsedID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	todo, err := h.todoService.CreateTodo(c.Request().Context(), &req, parsedID)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	response := models.TodoResponse{
		ID:todo.ID,
		Title: todo.Title,
		Description: todo.Description,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *TodoHandler) GetTodo(c echo.Context) error {

	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": "failed to get authenticated user",
			},
		)
	}
	parsedID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	todoID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": "invalid todo id",
			},
		)
	}

	todo, err := h.todoService.GetTodo(c.Request().Context(), parsedID, todoID)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	response := models.TodoResponse{
		ID:todo.ID,
		Title: todo.Title,
		Description: todo.Description,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
	return c.JSON(http.StatusOK, response)
}


func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	var req models.UpdateTodoRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	if err := c.Validate(&req); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)

		if !ok {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		var messages []string

		for _, validationErr := range validationErrors {
			switch validationErr.Tag() {
				case "required":
					messages = append(messages, validationErr.Field()+" is required.")
				case "notblank":
					messages = append(messages, validationErr.Field()+" should not be blank.")
			}
		}
		return c.JSON(
			http.StatusBadRequest,
			map[string]any{
				"errors": messages,
			},
		)
	}
	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": "failed to get authenticated user",
			},
		)
	}

	parsedID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	todoID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": "invalid todo id",
			},
		)
	}

	todo, err := h.todoService.UpdateTodo(c.Request().Context(), &req, todoID, parsedID)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	response := models.TodoResponse{
		ID:todo.ID,
		Title: todo.Title,
		Description: todo.Description,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *TodoHandler) ListTodos(c echo.Context) error {

	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": "failed to get authenticated user",
			},
		)
	}
	parsedID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	todos, err := h.todoService.ListTodos(c.Request().Context(), parsedID)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
	}
	return c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) DeleteTodo(c echo.Context) error {

	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": "failed to get authenticated user",
			},
		)
	}
	parsedID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	todoID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": "invalid todo id",
			},
		)
	}

	err = h.todoService.DeleteTodo(c.Request().Context(), todoID, parsedID)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(http.StatusNoContent, err)
}
