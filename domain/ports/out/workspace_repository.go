package ports_out

import domain "github.com/geordym/pendientico/domain/model"

type WorkspaceRepository interface {
	SaveWorkspace(workspace domain.Workspace) error
}
