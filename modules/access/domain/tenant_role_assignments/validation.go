package tenant_role_assignments

import (
	accessErr "nfxidentity/errors/src/access"

	"github.com/google/uuid"
)

func (a *TenantRoleAssignment) Validate() error {
	if a.UserID() == uuid.Nil || a.TenantID() == uuid.Nil || a.TenantRoleID() == uuid.Nil {
		return accessErr.ErrTenantRoleAssignmentNotFound
	}
	return nil
}
