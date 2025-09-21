package application_factory

import (
	usecase "github.com/geordym/pendientico/application/usecases"
	ports "github.com/geordym/pendientico/domain/ports/out"
)

func NewCreateUserUseCase(userRepository ports.UserRepository, authenticationProviderCommunication ports.AuthenticationProviderCommunication) *usecase.CreateUserUseCase {
	return usecase.NewCreateUser(userRepository, authenticationProviderCommunication)
}
