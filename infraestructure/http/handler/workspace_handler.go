package handler

import (
	"net/http"

	usecase "github.com/geordym/pendientico/application/usecases/workspaces"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go/log"
)

type WorkspaceHandler struct {
	createWorkspaceUseCase        usecase.CreateWorkspaceUseCase
	createWorkspaceContactUseCase usecase.CreateWorkspaceContactUseCase
}

func NewWorkspaceHandler(createWorkspaceUseCase usecase.CreateWorkspaceUseCase, createWorkspaceContactUseCase usecase.CreateWorkspaceContactUseCase) *WorkspaceHandler {
	return &WorkspaceHandler{createWorkspaceUseCase: createWorkspaceUseCase,
		createWorkspaceContactUseCase: createWorkspaceContactUseCase,
	}
}

func (h *WorkspaceHandler) HandleCreateWorkspace(c echo.Context) error {
	createWorkspaceCommand := usecase.CreateWorkspaceCommand{}

	if err := c.Bind(&createWorkspaceCommand); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request payload",
		})
	}

	err := h.createWorkspaceUseCase.Execute(c.Request().Context(), createWorkspaceCommand)
	if err != nil {
		log.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "workspace received",
	})
}

func (h *WorkspaceHandler) HandleCreateWorkspaceContact(c echo.Context) error {
	createWorkspaceContactCommand := usecase.CreateWorkspaceContactCommand{}

	if err := c.Bind(&createWorkspaceContactCommand); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request payload",
		})
	}

    workspaceID := c.Param("workspaceId")
	createWorkspaceContactCommand.WorkspaceID = workspaceID


	err := h.createWorkspaceContactUseCase.Execute(c.Request().Context(), createWorkspaceContactCommand)
	if err != nil {
		log.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "workspace contact received",
	})
}
