package routes

import (
	"github.com/HR-Shekhar/todo-api/internal/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(
	e *echo.Echo,
	userHandler *handler.UserHandler,
) {
	e.GET("/health", handler.HealthCheck)
	e.POST("/test-user", userHandler.CreateUser)
}
