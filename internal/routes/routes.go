package routes

import (
	"github.com/HR-Shekhar/todo-api/internal/handler"
	"github.com/HR-Shekhar/todo-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(
	e *echo.Echo,
	userHandler *handler.UserHandler,
	todoHandler *handler.TodoHandler,
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
	e.GET("/todos/:id", todoHandler.GetTodo, middleware.JWTMiddleware(jwtSecret))
	e.POST("/todos", todoHandler.CreateTodo, middleware.JWTMiddleware(jwtSecret))
	e.PUT("/todos/:id", todoHandler.UpdateTodo, middleware.JWTMiddleware(jwtSecret))
	e.GET("/todos", todoHandler.ListTodos, middleware.JWTMiddleware(jwtSecret))
	e.DELETE("/todos/:id", todoHandler.DeleteTodo, middleware.JWTMiddleware(jwtSecret))
}
