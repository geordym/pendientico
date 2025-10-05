package handler

import (
	"net/http"

	usecase "github.com/geordym/pendientico/application/usecases/users"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	createUserUseCase usecase.CreateUserUseCase
}

func NewUserHandler(createUserUseCase usecase.CreateUserUseCase) *UserHandler {
	return &UserHandler{createUserUseCase: createUserUseCase}
}

func (h *UserHandler) HandleCreateUser(c echo.Context) error {
	createUserCommand := usecase.CreateUserCommand{}

	if err := c.Bind(&createUserCommand); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request payload",
		})
	}

	h.createUserUseCase.Execute(createUserCommand)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user received",
	})
}
