package infraestructure

import (
	"github.com/geordym/pendientico/domain/model"
	ports "github.com/geordym/pendientico/domain/ports/out"
	"gorm.io/gorm"
)

type PostgresWorkspaceRepository struct {
	db *gorm.DB
}

func NewPostgresWorkspaceRepository(db *gorm.DB) *PostgresWorkspaceRepository {
	return &PostgresWorkspaceRepository{db: db}
}

func (w *PostgresWorkspaceRepository) SaveWorkspace(workspace model.Workspace) error {
	return w.db.Create(&workspace).Error
}

var _ ports.WorkspaceRepository = (*PostgresWorkspaceRepository)(nil)
