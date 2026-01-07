package respdto

import (
	"time"

	badgeAppViews "nfxid/modules/auth/application/badge/views"
	profileBadgeAppViews "nfxid/modules/auth/application/profile_badge/views"

	"github.com/google/uuid"
)

type ProfileBadgeDTO struct {
	ID          uuid.UUID `json:"id"`
	ProfileID   uuid.UUID `json:"profile_id"`
	BadgeID     uuid.UUID `json:"badge_id"`
	Description *string   `json:"description,omitempty"`
	Level       *int      `json:"level,omitempty"`
	EarnedAt    time.Time `json:"earned_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// Relations
	Badge *BadgeDTO `json:"badge,omitempty"` // 嵌套的徽章信息
}

// ProfileBadgeViewToDTO converts application ProfileBadgeView to response DTO
func ProfileBadgeViewToDTO(v *profileBadgeAppViews.ProfileBadgeView, badgeView *badgeAppViews.BadgeView) *ProfileBadgeDTO {
	if v == nil {
		return nil
	}

	dto := &ProfileBadgeDTO{
		ID:          v.ID,
		ProfileID:   v.ProfileID,
		BadgeID:     v.BadgeID,
		Description: v.Description,
		Level:       v.Level,
		EarnedAt:    v.EarnedAt,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}

	// 如果有嵌套的 Badge 信息，也转换
	if badgeView != nil {
		dto.Badge = BadgeViewToDTO(badgeView)
	}

	return dto
}

// ProfileBadgeListViewToDTO converts list of ProfileBadgeView to DTOs
func ProfileBadgeListViewToDTO(views []profileBadgeAppViews.ProfileBadgeView) []ProfileBadgeDTO {
	dtos := make([]ProfileBadgeDTO, len(views))
	for i, v := range views {
		if dto := ProfileBadgeViewToDTO(&v, nil); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
