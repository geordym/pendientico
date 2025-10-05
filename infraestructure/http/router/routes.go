package routes

import (
	"github.com/geordym/pendientico/infraestructure/http/handler"
	"github.com/geordym/pendientico/infraestructure/middlewares"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, userHandler *handler.UserHandler, workspaceHandler *handler.WorkspaceHandler) {
	//e.GET("/users", handlers.GetUsers, middlewares.AuthMiddleware)
	e.POST("/api/v1/users", userHandler.HandleCreateUser)
	e.POST("/api/v1/workspaces", workspaceHandler.HandleCreateWorkspace, middlewares.AuthMiddleware)
	e.POST("/api/v1/workspaces/:workspaceId/contacts", workspaceHandler.HandleCreateWorkspaceContact, middlewares.AuthMiddleware)

}
