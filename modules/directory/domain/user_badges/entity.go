package user_badges

import (
	"time"

	"github.com/google/uuid"
)

type UserBadge struct {
	state UserBadgeState
}

type UserBadgeState struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	BadgeID     uuid.UUID
	Description string
	Level       int
	EarnedAt    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (ub *UserBadge) ID() uuid.UUID        { return ub.state.ID }
func (ub *UserBadge) UserID() uuid.UUID    { return ub.state.UserID }
func (ub *UserBadge) BadgeID() uuid.UUID   { return ub.state.BadgeID }
func (ub *UserBadge) Description() string  { return ub.state.Description }
func (ub *UserBadge) Level() int           { return ub.state.Level }
func (ub *UserBadge) EarnedAt() time.Time  { return ub.state.EarnedAt }
func (ub *UserBadge) CreatedAt() time.Time { return ub.state.CreatedAt }
func (ub *UserBadge) UpdatedAt() time.Time { return ub.state.UpdatedAt }
