package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users := []string{"Juan", "Maria", "Pedro"}
	return c.JSON(http.StatusOK, users)
}
