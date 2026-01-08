package member_roles

import "github.com/google/uuid"

func (mr *MemberRole) Validate() error {
	if mr.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	if mr.MemberID() == uuid.Nil {
		return ErrMemberIDRequired
	}
	if mr.RoleID() == uuid.Nil {
		return ErrRoleIDRequired
	}
	return nil
}
