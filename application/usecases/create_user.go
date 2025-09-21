package usecase

import (
	"github.com/geordym/pendientico/domain/model"
	ports "github.com/geordym/pendientico/domain/ports/out"
	"github.com/labstack/gommon/log"
)

type CreateUserUseCase struct {
	userRepository                      ports.UserRepository
	authenticationProviderCommunication ports.AuthenticationProviderCommunication
}

func NewCreateUser(userRepository ports.UserRepository, authenticationProviderCommunication ports.AuthenticationProviderCommunication) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: userRepository,
		authenticationProviderCommunication: authenticationProviderCommunication}
}

type CreateUserCommand struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (uc *CreateUserUseCase) Execute(cmd CreateUserCommand) error {

	authProviderUserId, err := uc.authenticationProviderCommunication.SaveUser(cmd.Email, cmd.Password)
	if err != nil {
		log.Error("Ocurrio un error al crear el usuarion en el proveedor de autenticacion")
		return err
	}

	user := model.NewUser(authProviderUserId, cmd.Name, cmd.Phone, cmd.Email)

	err = uc.userRepository.SaveUser(user)
	if err != nil {
		return err
	}

	return nil
}
