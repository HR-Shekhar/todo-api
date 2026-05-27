package main

import (
	"github.com/HR-Shekhar/todo-api/internal/config"
	"github.com/HR-Shekhar/todo-api/internal/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.LoadConfig()
	routes.RegisterRoutes(e)
	
	if err := e.Start(":"+ cfg.Port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}