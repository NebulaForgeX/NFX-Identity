package results

import (
	"time"

	"nfxid/modules/directory/domain/user_badges"

	"github.com/google/uuid"
)

type UserBadgeRO struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	BadgeID     uuid.UUID
	Description string
	Level       int
	EarnedAt    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// UserBadgeMapper 将 Domain UserBadge 转换为 Application UserBadgeRO
func UserBadgeMapper(ub *user_badges.UserBadge) UserBadgeRO {
	if ub == nil {
		return UserBadgeRO{}
	}

	return UserBadgeRO{
		ID:          ub.ID(),
		UserID:      ub.UserID(),
		BadgeID:     ub.BadgeID(),
		Description: ub.Description(),
		Level:       ub.Level(),
		EarnedAt:    ub.EarnedAt(),
		CreatedAt:   ub.CreatedAt(),
		UpdatedAt:   ub.UpdatedAt(),
	}
}
