package member_groups

import (
	tenantsErr "nfxid/errors/src/tenants"

	"github.com/google/uuid"
)

func (mg *MemberGroup) Validate() error {
	if mg.MemberID() == uuid.Nil {
		return tenantsErr.ErrMemberIDRequired
	}
	if mg.GroupID() == uuid.Nil {
		return tenantsErr.ErrGroupIDRequired
	}
	return nil
}
