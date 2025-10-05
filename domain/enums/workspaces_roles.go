package domain 

type WorkspaceRoleEnum string

const (
	WorkspaceRoleOwner       WorkspaceRoleEnum = "OWNER"
	WorkspaceRoleCollaborator WorkspaceRoleEnum = "COLLABORATOR"
)
