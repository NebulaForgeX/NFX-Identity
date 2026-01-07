package profile_badge

import (
	"time"

	profileBadgeErrors "nfxid/modules/auth/domain/profile_badge/errors"

	"github.com/google/uuid"
)

type NewProfileBadgeParams struct {
	ProfileID uuid.UUID
	BadgeID   uuid.UUID
	Editable  ProfileBadgeEditable
}

func NewProfileBadge(p NewProfileBadgeParams) (*ProfileBadge, error) {
	if p.ProfileID == uuid.Nil {
		return nil, profileBadgeErrors.ErrProfileIDRequired
	}
	if p.BadgeID == uuid.Nil {
		return nil, profileBadgeErrors.ErrBadgeIDRequired
	}
	if err := p.Editable.Validate(); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewProfileBadgeFromState(ProfileBadgeState{
		ID:        id,
		ProfileID: p.ProfileID,
		BadgeID:   p.BadgeID,
		Editable:  p.Editable,
		EarnedAt:  now,
		CreatedAt: now,
		UpdatedAt: now,
	}), nil
}

func NewProfileBadgeFromState(st ProfileBadgeState) *ProfileBadge {
	return &ProfileBadge{state: st}
}
