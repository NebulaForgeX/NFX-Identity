package member_app_roles

import "github.com/google/uuid"

func (mar *MemberAppRole) Validate() error {
	if mar.MemberID() == uuid.Nil {
		return ErrMemberIDRequired
	}
	if mar.AppID() == uuid.Nil {
		return ErrAppIDRequired
	}
	if mar.RoleID() == uuid.Nil {
		return ErrRoleIDRequired
	}
	return nil
}
