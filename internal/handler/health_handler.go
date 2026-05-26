package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Health struct {
	Status string `json:"status"`
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, Health{Status: "OK"})
}
