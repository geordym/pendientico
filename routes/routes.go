package routes

import (
	"github.com/geordym/pendientico/application/handlers"
	"github.com/geordym/pendientico/middlewares"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.GET("/users", handlers.GetUsers, middlewares.AuthMiddleware)
}
