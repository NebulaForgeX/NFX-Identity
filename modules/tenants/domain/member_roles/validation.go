package member_roles

import (
	tenantsErr "nfxidentity/errors/src/tenants"

	"github.com/google/uuid"
)

func (mr *MemberRole) Validate() error {
	if mr.TenantID() == uuid.Nil {
		return tenantsErr.ErrTenantIDRequired
	}
	if mr.MemberID() == uuid.Nil {
		return tenantsErr.ErrMemberIDRequired
	}
	if mr.RoleID() == uuid.Nil {
		return tenantsErr.ErrRoleIDRequired
	}
	return nil
}
