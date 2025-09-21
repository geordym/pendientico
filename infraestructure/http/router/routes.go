package routes

import (
	"github.com/geordym/pendientico/infraestructure/http/handler"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, userHandler *handler.UserHandler) {
	//e.GET("/users", handlers.GetUsers, middlewares.AuthMiddleware)
	e.POST("/api/v1/users", userHandler.HandleCreateUser)
}
