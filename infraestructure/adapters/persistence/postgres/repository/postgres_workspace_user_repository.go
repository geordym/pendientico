package infraestructure

import (
	"github.com/geordym/pendientico/domain/model"
	ports_out "github.com/geordym/pendientico/domain/ports/out"
	"gorm.io/gorm"
)

type PostgresWorkspaceMembersRepository struct {
	db *gorm.DB
}

func NewPostgresWorkspaceMembersRepository(db *gorm.DB) *PostgresWorkspaceMembersRepository {
	return &PostgresWorkspaceMembersRepository{db: db}
}

func (wu *PostgresWorkspaceMembersRepository) SaveWorkSpaceMember(workspaceUser model.WorkspaceMember) error {
	return wu.db.Create(&workspaceUser).Error
}

var _ ports_out.WorkspaceMembersRepository = (*PostgresWorkspaceMembersRepository)(nil)
