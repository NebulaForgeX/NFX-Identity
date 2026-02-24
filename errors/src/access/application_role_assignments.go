package access

import "nfxid/pkgs/errx"

const (
	CodeApplicationRoleAssignmentNotFound = "APPLICATION_ROLE_ASSIGNMENT_NOT_FOUND"
)

var (
	ErrApplicationRoleAssignmentNotFound = errx.NotFound(CodeApplicationRoleAssignmentNotFound, "application role assignment not found")
)
