package usecase

import (
	"context"

	domain "github.com/geordym/pendientico/domain/enums"
	"github.com/geordym/pendientico/domain/model"
	ports_out "github.com/geordym/pendientico/domain/ports/out"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

type CreateWorkspaceUseCase struct {
	workspaceRepository                 ports_out.WorkspaceRepository
	workspaceUsersRepository            ports_out.WorkspaceMembersRepository
	userRepository                      ports_out.UserRepository
	authenticationProviderCommunication ports_out.AuthenticationProviderCommunication
}

func NewCreateWorkspace(workspaceRepository ports_out.WorkspaceRepository,
	workspaceUsersRepository ports_out.WorkspaceMembersRepository,
	userRepository ports_out.UserRepository,
	authenticationProviderCommunication ports_out.AuthenticationProviderCommunication,
) *CreateWorkspaceUseCase {
	return &CreateWorkspaceUseCase{workspaceRepository: workspaceRepository,
		workspaceUsersRepository:            workspaceUsersRepository,
		userRepository:                      userRepository,
		authenticationProviderCommunication: authenticationProviderCommunication,
	}

}

func (uc *CreateWorkspaceUseCase) Execute(ctx context.Context, cmd CreateWorkspaceCommand) error {

	workspace := model.Workspace{
		ID:   uuid.NewString(),
		Name: cmd.Name,
	}

	err := uc.workspaceRepository.SaveWorkspace(workspace)
	if err != nil {
		log.Error("Ocurrio un error al intentar crear el workspace")
		return err
	}

	authProviderUserId, err := uc.authenticationProviderCommunication.GetUserAuthProviderIdLogged(ctx)
	if err != nil {
		log.Error("Ocurrio un error al extraer el userAuthProviderId del context")
		return err
	}

	user, err := uc.userRepository.FindUserByAuthProviderUserId(authProviderUserId)
	if err != nil {
		log.Error("Ocurrio un error al buscar el usuario en la BD por el AuthProviderUserId")
	}

	workspaceUserOwner := model.WorkspaceMember{
		ID:          uuid.NewString(),
		WorkspaceId: workspace.ID,
		UserId:      user.ID,
		Role:        string(domain.WorkspaceRoleOwner),
	}

	err = uc.workspaceUsersRepository.SaveWorkSpaceMember(workspaceUserOwner)
	if err != nil {
		log.Error("ocurrio un error al guardar el workpsace user owner")
		return err
	}

	return nil
}

type CreateWorkspaceCommand struct {
	Name string `json:"name"`
}
