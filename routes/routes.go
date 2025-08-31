package routes

import (
	"github.com/geordym/pendientico/handlers"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.GET("/users", handlers.GetUsers)
}
