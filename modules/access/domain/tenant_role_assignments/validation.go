package tenant_role_assignments

import "github.com/google/uuid"

func (a *TenantRoleAssignment) Validate() error {
	if a.UserID() == uuid.Nil || a.TenantID() == uuid.Nil || a.TenantRoleID() == uuid.Nil {
		return ErrTenantRoleAssignmentNotFound
	}
	return nil
}
