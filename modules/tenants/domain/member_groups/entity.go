package member_groups

import (
	"time"

	"github.com/google/uuid"
)

type MemberGroup struct {
	state MemberGroupState
}

type MemberGroupState struct {
	ID         uuid.UUID
	MemberID   uuid.UUID
	GroupID    uuid.UUID
	AssignedAt time.Time
	AssignedBy *uuid.UUID
	RevokedAt  *time.Time
	RevokedBy  *uuid.UUID
}

func (mg *MemberGroup) ID() uuid.UUID        { return mg.state.ID }
func (mg *MemberGroup) MemberID() uuid.UUID   { return mg.state.MemberID }
func (mg *MemberGroup) GroupID() uuid.UUID    { return mg.state.GroupID }
func (mg *MemberGroup) AssignedAt() time.Time { return mg.state.AssignedAt }
func (mg *MemberGroup) AssignedBy() *uuid.UUID { return mg.state.AssignedBy }
func (mg *MemberGroup) RevokedAt() *time.Time { return mg.state.RevokedAt }
func (mg *MemberGroup) RevokedBy() *uuid.UUID { return mg.state.RevokedBy }
