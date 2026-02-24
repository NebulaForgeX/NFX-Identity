package access

import "nfxid/pkgs/errx"

const (
	CodeTenantRoleAssignmentNotFound = "TENANT_ROLE_ASSIGNMENT_NOT_FOUND"
)

var (
	ErrTenantRoleAssignmentNotFound = errx.NotFound(CodeTenantRoleAssignmentNotFound, "tenant role assignment not found")
)

/*
!TENANT_ROLE_ASSIGNMENT_NOT_FOUND
*en<tenant role assignment not found>
*zh<租户角色分配不存在>
*fr<affectation de rôle tenant introuvable>

*/
