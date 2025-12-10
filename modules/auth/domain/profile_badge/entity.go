package profile_badge

import (
	"time"

	"github.com/google/uuid"
)

type ProfileBadge struct {
	state ProfileBadgeState
}

type ProfileBadgeState struct {
	ID          uuid.UUID
	ProfileID   uuid.UUID
	BadgeID     uuid.UUID
	Editable    ProfileBadgeEditable
	EarnedAt    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProfileBadgeEditable struct {
	Description *string
	Level       *int
}

func (pb *ProfileBadge) ID() uuid.UUID                    { return pb.state.ID }
func (pb *ProfileBadge) ProfileID() uuid.UUID              { return pb.state.ProfileID }
func (pb *ProfileBadge) BadgeID() uuid.UUID                { return pb.state.BadgeID }
func (pb *ProfileBadge) Editable() ProfileBadgeEditable   { return pb.state.Editable }
func (pb *ProfileBadge) EarnedAt() time.Time              { return pb.state.EarnedAt }
func (pb *ProfileBadge) CreatedAt() time.Time             { return pb.state.CreatedAt }
func (pb *ProfileBadge) UpdatedAt() time.Time             { return pb.state.UpdatedAt }
