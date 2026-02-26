package member_app_roles

import (
	tenantsErr "nfxidentity/errors/src/tenants"

	"github.com/google/uuid"
)

func (mar *MemberAppRole) Validate() error {
	if mar.MemberID() == uuid.Nil {
		return tenantsErr.ErrMemberIDRequired
	}
	if mar.AppID() == uuid.Nil {
		return tenantsErr.ErrAppIDRequired
	}
	if mar.RoleID() == uuid.Nil {
		return tenantsErr.ErrRoleIDRequired
	}
	return nil
}
