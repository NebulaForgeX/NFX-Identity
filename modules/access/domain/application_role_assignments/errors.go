package application_role_assignments

import "nfxid/pkgs/errx"

var (
	ErrApplicationRoleAssignmentNotFound = errx.NotFound("APPLICATION_ROLE_ASSIGNMENT_NOT_FOUND", "application role assignment not found")
)
