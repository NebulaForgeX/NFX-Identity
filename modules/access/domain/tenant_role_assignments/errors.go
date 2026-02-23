package tenant_role_assignments

import "nfxid/pkgs/errx"

var (
	ErrTenantRoleAssignmentNotFound = errx.NotFound("TENANT_ROLE_ASSIGNMENT_NOT_FOUND", "tenant role assignment not found")
)
