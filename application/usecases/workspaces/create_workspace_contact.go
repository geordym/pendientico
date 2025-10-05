package usecase

import (
	"context"

	"github.com/geordym/pendientico/domain/model"
	ports_out "github.com/geordym/pendientico/domain/ports/out"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

type CreateWorkspaceContactUseCase struct {
	workspaceContactRepository ports_out.WorkspaceContactRepository
}

func NewCreateWorkspaceContactUseCase(workspaceContactRepository ports_out.WorkspaceContactRepository) *CreateWorkspaceContactUseCase {
	return &CreateWorkspaceContactUseCase{workspaceContactRepository: workspaceContactRepository}
}

func (uc *CreateWorkspaceContactUseCase) Execute(ctx context.Context, cmd CreateWorkspaceContactCommand) error {

	workspaceContact := model.WorkspaceContact{
		ID:          uuid.NewString(),
		WorkspaceID: cmd.WorkspaceID,
		Name:        cmd.Name,
	}

	err := uc.workspaceContactRepository.SaveWorkspaceContact(workspaceContact)
	if err != nil {
		log.Error("Ocurrio un error al guardar el contacto en el workspace")
		return err
	}

	return nil
}

type CreateWorkspaceContactCommand struct {
	WorkspaceID string `json:"workspaceId"`
	Name        string `json:"name"`
}
