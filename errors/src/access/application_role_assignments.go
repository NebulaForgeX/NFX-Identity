package access

import "nfxidentity/pkgs/errx"

const (
	CodeApplicationRoleAssignmentNotFound = "APPLICATION_ROLE_ASSIGNMENT_NOT_FOUND"
)

var (
	ErrApplicationRoleAssignmentNotFound = errx.NotFound(CodeApplicationRoleAssignmentNotFound, "application role assignment not found")
)

/*
!APPLICATION_ROLE_ASSIGNMENT_NOT_FOUND
*en<application role assignment not found>
*zh<应用角色分配不存在>
*fr<application role assignment non trouvée>
*/