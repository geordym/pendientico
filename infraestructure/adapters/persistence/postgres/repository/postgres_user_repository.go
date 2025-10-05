// infra/postgres_user_repository.go
package infraestructure

import (
	"github.com/geordym/pendientico/domain/model"
	ports "github.com/geordym/pendientico/domain/ports/out"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) SaveUser(user model.User) error {
	return r.db.Create(&user).Error
}

func (r *PostgresUserRepository) FindUserByAuthProviderUserId(authProviderUserId string) (*model.User, error) {
	var user model.User
	err := r.db.Where("auth_provider_user_id = ?", authProviderUserId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

var _ ports.UserRepository = (*PostgresUserRepository)(nil)
