package views

import (
	"time"

	profileBadgeDomainViews "nfxid/modules/auth/domain/profile_badge/views"
	userDomainViews "nfxid/modules/auth/domain/user/views"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type ProfileBadgeView struct {
	ID          uuid.UUID `json:"id"`
	ProfileID   uuid.UUID `json:"profile_id"`
	BadgeID     uuid.UUID `json:"badge_id"`
	Description *string   `json:"description"`
	Level       *int      `json:"level"`
	EarnedAt    time.Time `json:"earned_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProfileBadgeViewMapper 将 Domain ProfileBadgeView 转换为 Application ProfileBadgeView
func ProfileBadgeViewMapper(v profileBadgeDomainViews.ProfileBadgeView) ProfileBadgeView {
	return ProfileBadgeView{
		ID:          v.ID,
		ProfileID:   v.ProfileID,
		BadgeID:     v.BadgeID,
		Description: v.Description,
		Level:       v.Level,
		EarnedAt:    v.EarnedAt,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}
}

type UserBadgeView struct {
	UserID      uuid.UUID       `json:"user_id"`
	Username    string          `json:"username"`
	Email       string          `json:"email"`
	ProfileID   *uuid.UUID      `json:"profile_id"`
	DisplayName *string         `json:"display_name"`
	Badges      *datatypes.JSON `json:"badges"`
}

// UserBadgeViewMapper 将 Domain UserBadgesView 转换为 Application UserBadgeView
func UserBadgeViewMapper(v userDomainViews.UserBadgesView) UserBadgeView {
	return UserBadgeView{
		UserID:      v.UserID,
		Username:    v.Username,
		Email:       v.Email,
		ProfileID:   v.ProfileID,
		DisplayName: v.DisplayName,
		Badges:      v.Badges,
	}
}
