package ports_out

import model "github.com/geordym/pendientico/domain/model"

type UserRepository interface {
	SaveUser(user model.User) error
	FindUserByAuthProviderUserId(authProviderUserId string) (*model.User, error)
}
