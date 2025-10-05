package application_factory

import (
	users_usecases "github.com/geordym/pendientico/application/usecases/users"
	workspace_usecases "github.com/geordym/pendientico/application/usecases/workspaces"
	ports "github.com/geordym/pendientico/domain/ports/out"
)

func NewCreateUserUseCase(userRepository ports.UserRepository, authenticationProviderCommunication ports.AuthenticationProviderCommunication) *users_usecases.CreateUserUseCase {
	return users_usecases.NewCreateUser(userRepository, authenticationProviderCommunication)
}

func NewCreateWorkspaceUseCase(workspaceRepository ports.WorkspaceRepository, workspaceUsersRepository ports.WorkspaceMembersRepository, userRepository ports.UserRepository, authProviderCommunication ports.AuthenticationProviderCommunication) *workspace_usecases.CreateWorkspaceUseCase {
	return workspace_usecases.NewCreateWorkspace(workspaceRepository, workspaceUsersRepository, userRepository, authProviderCommunication)
}

func NewCreateWorkspaceContactUseCase(workspaceContactRepository ports.WorkspaceContactRepository) *workspace_usecases.CreateWorkspaceContactUseCase {
	return workspace_usecases.NewCreateWorkspaceContactUseCase(workspaceContactRepository)
}
