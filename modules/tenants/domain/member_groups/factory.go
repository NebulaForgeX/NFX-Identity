package member_groups

import (
	"time"

	"github.com/google/uuid"
)

type NewMemberGroupParams struct {
	MemberID   uuid.UUID
	GroupID    uuid.UUID
	AssignedBy *uuid.UUID
}

func NewMemberGroup(p NewMemberGroupParams) (*MemberGroup, error) {
	if err := validateMemberGroupParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewMemberGroupFromState(MemberGroupState{
		ID:         id,
		MemberID:   p.MemberID,
		GroupID:    p.GroupID,
		AssignedAt: now,
		AssignedBy: p.AssignedBy,
	}), nil
}

func NewMemberGroupFromState(st MemberGroupState) *MemberGroup {
	return &MemberGroup{state: st}
}

func validateMemberGroupParams(p NewMemberGroupParams) error {
	if p.MemberID == uuid.Nil {
		return ErrMemberIDRequired
	}
	if p.GroupID == uuid.Nil {
		return ErrGroupIDRequired
	}
	return nil
}
