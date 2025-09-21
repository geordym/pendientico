package postgres_factory

import (
	ports "github.com/geordym/pendientico/domain/ports/out"
	infraestructure "github.com/geordym/pendientico/infraestructure/adapters/persistence/postgres/repository"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return infraestructure.NewPostgresUserRepository(db)
}