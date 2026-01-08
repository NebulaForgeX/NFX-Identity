package member_app_roles

import (
	"time"

	"github.com/google/uuid"
)

type NewMemberAppRoleParams struct {
	MemberID   uuid.UUID
	AppID      uuid.UUID
	RoleID     uuid.UUID
	AssignedBy *uuid.UUID
	ExpiresAt  *time.Time
}

func NewMemberAppRole(p NewMemberAppRoleParams) (*MemberAppRole, error) {
	if err := validateMemberAppRoleParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewMemberAppRoleFromState(MemberAppRoleState{
		ID:         id,
		MemberID:   p.MemberID,
		AppID:      p.AppID,
		RoleID:     p.RoleID,
		AssignedAt: now,
		AssignedBy: p.AssignedBy,
		ExpiresAt:  p.ExpiresAt,
	}), nil
}

func NewMemberAppRoleFromState(st MemberAppRoleState) *MemberAppRole {
	return &MemberAppRole{state: st}
}

func validateMemberAppRoleParams(p NewMemberAppRoleParams) error {
	if p.MemberID == uuid.Nil {
		return ErrMemberIDRequired
	}
	if p.AppID == uuid.Nil {
		return ErrAppIDRequired
	}
	if p.RoleID == uuid.Nil {
		return ErrRoleIDRequired
	}
	return nil
}
