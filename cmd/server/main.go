package main

import (
	"github.com/HR-Shekhar/todo-api/internal/config"
	"github.com/HR-Shekhar/todo-api/internal/database"
	"github.com/HR-Shekhar/todo-api/internal/handler"
	"github.com/HR-Shekhar/todo-api/internal/repository"
	"github.com/HR-Shekhar/todo-api/internal/routes"
	"github.com/HR-Shekhar/todo-api/internal/service"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()
	
	e := echo.New()

	dbpool := database.NewPostgresConnection(cfg)
	userRepo := repository.NewUserRepository(dbpool)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)
	defer dbpool.Close()

	routes.RegisterRoutes(
		e,
		userHandler,
	)
	
	if err := e.Start(":"+ cfg.Port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}