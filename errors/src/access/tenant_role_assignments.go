package access

import "nfxid/pkgs/errx"

const (
	CodeTenantRoleAssignmentNotFound = "TENANT_ROLE_ASSIGNMENT_NOT_FOUND"
)

var (
	ErrTenantRoleAssignmentNotFound = errx.NotFound(CodeTenantRoleAssignmentNotFound, "tenant role assignment not found")
)
