package reqdto

import (
	userBadgeAppCommands "nfxid/modules/directory/application/user_badges/commands"

	"github.com/google/uuid"
)

type UserBadgeCreateRequestDTO struct {
	UserID      uuid.UUID `json:"user_id" validate:"required,uuid"`
	BadgeID     uuid.UUID `json:"badge_id" validate:"required,uuid"`
	Description string    `json:"description"`
	Level       int       `json:"level"`
	EarnedAt    *string   `json:"earned_at,omitempty"`
}

func (r *UserBadgeCreateRequestDTO) ToCreateCmd() userBadgeAppCommands.CreateUserBadgeCmd {
	return userBadgeAppCommands.CreateUserBadgeCmd{
		UserID:      r.UserID,
		BadgeID:     r.BadgeID,
		Description: r.Description,
		Level:       r.Level,
		EarnedAt:    r.EarnedAt,
	}
}
