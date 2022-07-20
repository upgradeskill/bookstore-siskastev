package middleware

import (
	"bookstore-siskastev/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func Auth(ctx echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		bearerToken := (strings.Split(authHeader, " "))[1]
		if bearerToken == "" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		_, err := helpers.ValidateToken(bearerToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return ctx(c)
	}
}
