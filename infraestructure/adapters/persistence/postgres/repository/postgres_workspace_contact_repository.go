package infraestructure

import (
	"github.com/geordym/pendientico/domain/model"
	ports_out "github.com/geordym/pendientico/domain/ports/out"
	"gorm.io/gorm"
)

type PostgresWorkspaceContactRepository struct {
	db *gorm.DB
}

func NewPostgresWorkspaceContactRepository(db *gorm.DB) *PostgresWorkspaceContactRepository {
	return &PostgresWorkspaceContactRepository{db: db}
}

func (p *PostgresWorkspaceContactRepository) SaveWorkspaceContact(workspaceContact model.WorkspaceContact) error {
	return p.db.Create(&workspaceContact).Error
}

var _ ports_out.WorkspaceContactRepository = (*PostgresWorkspaceContactRepository)(nil)
