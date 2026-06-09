package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(
					http.StatusUnauthorized,
					map[string]string{
						"error": "missing authorization header",
					},
				)
			}

			parts := strings.Split(authHeader, " ")

			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(
					http.StatusUnauthorized,
					map[string]string{
						"error": "invalid authorization header",
					},
				)
			}
			tokenString := parts[1]
			token, err := jwt.Parse(
				tokenString,
				func(token *jwt.Token) (any, error) {
					return []byte(jwtSecret), nil
				},
				jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
			)
			if err != nil || !token.Valid {
				return c.JSON(
					http.StatusUnauthorized,
					map[string]string{
						"error": "invalid token",
					},
				)
			}
			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok {
				return c.JSON(
					http.StatusUnauthorized,
					map[string]string{
						"error": "invalid token claims",
					},
				)
			}

			userID, ok := claims["sub"].(string)

			if !ok {
				return c.JSON(
					http.StatusUnauthorized,
					map[string]string{
						"error": "invalid subject claim",
					},
				)
			}

			c.Set("userID", userID)

			return next(c)
		}
	}
}