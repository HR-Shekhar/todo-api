package routes

import (
	"github.com/HR-Shekhar/todo-api/internal/handler"
	"github.com/HR-Shekhar/todo-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(
	e *echo.Echo,
	userHandler *handler.UserHandler,
	jwtSecret string,
) {
	e.GET("/health", handler.HealthCheck)
	e.POST("/test-user", userHandler.CreateUser)
	e.POST("/login", userHandler.LoginUser)
	e.GET(
		"/me",
		userHandler.GetMe,
		middleware.JWTMiddleware(jwtSecret),
	)
}
