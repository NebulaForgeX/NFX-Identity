package views

import (
	"time"

	"github.com/google/uuid"
)

type ProfileBadgeView struct {
	ID          uuid.UUID
	ProfileID   uuid.UUID
	BadgeID     uuid.UUID
	Description *string
	Level       *int
	EarnedAt    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

