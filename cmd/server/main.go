package main

import (
	"github.com/HR-Shekhar/todo-api/internal/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/health", handler.HealthCheck)

	if err := e.Start(":8000"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}