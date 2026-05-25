package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Health struct{
	Status string `json:"status"`
}

func main() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Health{Status: "ok"})
	})

	if err := e.Start(":8000"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}