package ports_out

import domain "github.com/geordym/pendientico/domain/model"

type WorkspaceMembersRepository interface {
	SaveWorkSpaceMember(workspaceUser domain.WorkspaceMember) error
}
