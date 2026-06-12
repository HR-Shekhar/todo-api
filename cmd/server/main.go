package main

import (
	"github.com/HR-Shekhar/todo-api/internal/config"
	"github.com/HR-Shekhar/todo-api/internal/database"
	"github.com/HR-Shekhar/todo-api/internal/handler"
	"github.com/HR-Shekhar/todo-api/internal/repository"
	"github.com/HR-Shekhar/todo-api/internal/routes"
	"github.com/HR-Shekhar/todo-api/internal/service"
	"github.com/HR-Shekhar/todo-api/internal/validator"
	
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()
	
	e := echo.New()

	e.Validator = validator.NewCustomValidator()

	dbpool := database.NewPostgresConnection(cfg)
	
	userRepo := repository.NewUserRepository(dbpool)
	todoRepo := repository.NewTodoReposiory(dbpool)

	userService := service.NewUserService(userRepo, cfg.JWTSecret)
	todoService := service.NewTodoService(todoRepo)

	userHandler := handler.NewUserHandler(userService)
	todoHandler := handler.NewTodoHandler(todoService)
	defer dbpool.Close()

	routes.RegisterRoutes(
		e,
		userHandler,
		todoHandler,
		cfg.JWTSecret,
	)
	
	if err := e.Start(":"+ cfg.Port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}