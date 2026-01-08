package member_roles

import (
	"time"

	"github.com/google/uuid"
)

type NewMemberRoleParams struct {
	TenantID   uuid.UUID
	MemberID   uuid.UUID
	RoleID     uuid.UUID
	AssignedBy *uuid.UUID
	ExpiresAt  *time.Time
	Scope      *string
}

func NewMemberRole(p NewMemberRoleParams) (*MemberRole, error) {
	if err := validateMemberRoleParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewMemberRoleFromState(MemberRoleState{
		ID:         id,
		TenantID:   p.TenantID,
		MemberID:   p.MemberID,
		RoleID:     p.RoleID,
		AssignedAt: now,
		AssignedBy: p.AssignedBy,
		ExpiresAt:  p.ExpiresAt,
		Scope:      p.Scope,
	}), nil
}

func NewMemberRoleFromState(st MemberRoleState) *MemberRole {
	return &MemberRole{state: st}
}

func validateMemberRoleParams(p NewMemberRoleParams) error {
	if p.TenantID == uuid.Nil {
		return ErrTenantIDRequired
	}
	if p.MemberID == uuid.Nil {
		return ErrMemberIDRequired
	}
	if p.RoleID == uuid.Nil {
		return ErrRoleIDRequired
	}
	return nil
}
