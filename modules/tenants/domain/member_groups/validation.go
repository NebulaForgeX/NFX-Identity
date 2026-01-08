package member_groups

import "github.com/google/uuid"

func (mg *MemberGroup) Validate() error {
	if mg.MemberID() == uuid.Nil {
		return ErrMemberIDRequired
	}
	if mg.GroupID() == uuid.Nil {
		return ErrGroupIDRequired
	}
	return nil
}
