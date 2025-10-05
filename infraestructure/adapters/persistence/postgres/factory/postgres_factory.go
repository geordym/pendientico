package postgres_factory

import (
	ports "github.com/geordym/pendientico/domain/ports/out"
	infraestructure "github.com/geordym/pendientico/infraestructure/adapters/persistence/postgres/repository"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return infraestructure.NewPostgresUserRepository(db)
}

func NewWorkspaceRepository(db *gorm.DB) ports.WorkspaceRepository {
	return infraestructure.NewPostgresWorkspaceRepository(db)
}

func NewWorkspaceUsersRepository(db *gorm.DB) ports.WorkspaceMembersRepository {
	return infraestructure.NewPostgresWorkspaceMembersRepository(db)
}

func NewWorkspaceContactRepository(db *gorm.DB) ports.WorkspaceContactRepository {
	return infraestructure.NewPostgresWorkspaceContactRepository(db)
}
