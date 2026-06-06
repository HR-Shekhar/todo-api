package handler

import (
	// "errors"
	"net/http"

	"github.com/HR-Shekhar/todo-api/internal/models"
	"github.com/HR-Shekhar/todo-api/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}


func (h *UserHandler) CreateUser(c echo.Context) error {
	var req models.CreateUserRequest
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

		for _ ,validationErr := range validationErrors {
			switch validationErr.Tag() {
			case "required":
				messages = append(messages, validationErr.Field() + " is required.")
			case "email":
				messages = append(messages, validationErr.Field() + " must be valid.")
			case "min":
				messages = append(messages, validationErr.Field() + " must be at least " + validationErr.Param() + " characters.")
			case "notblank":
				messages = append(messages, validationErr.Field() + " should not be blank.")
			case "strongpassword":
				messages = append(messages, "password should include atleast an uppercase letter, a lowercase letter, a digit and a special character, spaces are not allowed.")
			}
		}
		return c.JSON(
			http.StatusBadRequest,
			map[string]any{
				"errors": messages,
			},
		)
	}

	user, err := h.userService.RegisterUser(
		c.Request().Context(),
		&req,
	)
	if err != nil {
		switch err {
			case service.ErrEmailAlreadyExists:
				return c.JSON(http.StatusConflict, map[string]string{
					"error": err.Error(),
				})
			case service.ErrUsernameAlreadyExists:
				return c.JSON(http.StatusConflict, map[string]string{
					"error": err.Error(),
				})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}
	response := models.UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return c.JSON(http.StatusCreated, response)
}