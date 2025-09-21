package ports_out

import domain "github.com/geordym/pendientico/domain/model"

type UserRepository interface {
	SaveUser(user domain.User) error
}
