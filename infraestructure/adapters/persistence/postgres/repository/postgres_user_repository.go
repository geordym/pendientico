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

var _ ports.UserRepository = (*PostgresUserRepository)(nil)
