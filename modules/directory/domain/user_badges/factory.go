package user_badges

import (
	"time"

	"github.com/google/uuid"
)

type NewUserBadgeParams struct {
	UserID      uuid.UUID
	BadgeID     uuid.UUID
	Description string
	Level       int
	EarnedAt    *time.Time
}

func NewUserBadge(p NewUserBadgeParams) (*UserBadge, error) {
	if err := validateUserBadgeParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	earnedAt := now
	if p.EarnedAt != nil {
		earnedAt = *p.EarnedAt
	}

	level := p.Level
	if level <= 0 {
		level = 1
	}

	return NewUserBadgeFromState(UserBadgeState{
		ID:          id,
		UserID:      p.UserID,
		BadgeID:     p.BadgeID,
		Description: p.Description,
		Level:       level,
		EarnedAt:    earnedAt,
		CreatedAt:   now,
		UpdatedAt:   now,
	}), nil
}

func NewUserBadgeFromState(st UserBadgeState) *UserBadge {
	return &UserBadge{state: st}
}

func validateUserBadgeParams(p NewUserBadgeParams) error {
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if p.BadgeID == uuid.Nil {
		return ErrBadgeIDRequired
	}
	return nil
}
