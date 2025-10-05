package ports_out

import "github.com/geordym/pendientico/domain/model"

type WorkspaceContactRepository interface {
	SaveWorkspaceContact(workspaceContact model.WorkspaceContact) error
}
