package handler

import (
	"errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func getAuthenticatedUserID(
	c echo.Context,
) (uuid.UUID, error) {

	userID, ok := c.Get("userID").(string)
	if !ok {
		return uuid.Nil, errors.New(
			"failed to get authenticated user",
		)
	}

	parsedID, err := uuid.Parse(userID)
	if err != nil {
		return uuid.Nil, err
	}

	return parsedID, nil
}