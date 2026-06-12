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
	e.GET("/todo/:id", todoHandler.GetTodo, middleware.JWTMiddleware(jwtSecret))
	e.POST("/todo", todoHandler.CreateTodo, middleware.JWTMiddleware(jwtSecret))
	e.PUT("/todo/:id", todoHandler.UpdateTodo, middleware.JWTMiddleware(jwtSecret))
	e.GET("/todo", todoHandler.ListTodos, middleware.JWTMiddleware(jwtSecret))
	e.DELETE("/todo", todoHandler.DeleteTodo, middleware.JWTMiddleware(jwtSecret))
}
