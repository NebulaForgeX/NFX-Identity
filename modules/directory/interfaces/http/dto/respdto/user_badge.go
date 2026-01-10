package respdto

import (
	"time"

	userBadgeAppResult "nfxid/modules/directory/application/user_badges/results"

	"github.com/google/uuid"
)

type UserBadgeDTO struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	BadgeID     uuid.UUID `json:"badge_id"`
	Description string    `json:"description"`
	Level       int       `json:"level"`
	EarnedAt    time.Time `json:"earned_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UserBadgeROToDTO converts application UserBadgeRO to response DTO
func UserBadgeROToDTO(v *userBadgeAppResult.UserBadgeRO) *UserBadgeDTO {
	if v == nil {
		return nil
	}

	return &UserBadgeDTO{
		ID:          v.ID,
		UserID:      v.UserID,
		BadgeID:     v.BadgeID,
		Description: v.Description,
		Level:       v.Level,
		EarnedAt:    v.EarnedAt,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}
}

// UserBadgeListROToDTO converts list of UserBadgeRO to DTOs
func UserBadgeListROToDTO(results []userBadgeAppResult.UserBadgeRO) []UserBadgeDTO {
	dtos := make([]UserBadgeDTO, len(results))
	for i, v := range results {
		if dto := UserBadgeROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
