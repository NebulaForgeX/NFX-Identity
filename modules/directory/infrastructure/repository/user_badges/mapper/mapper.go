package mapper

import (
	"nfxid/modules/directory/domain/user_badges"
	"nfxid/modules/directory/infrastructure/rdb/models"
)

// UserBadgeDomainToModel 将 Domain UserBadge 转换为 Model UserBadge
func UserBadgeDomainToModel(ub *user_badges.UserBadge) *models.UserBadge {
	if ub == nil {
		return nil
	}

	var description *string
	if ub.Description() != "" {
		d := ub.Description()
		description = &d
	}

	level := ub.Level()

	return &models.UserBadge{
		ID:          ub.ID(),
		UserID:      ub.UserID(),
		BadgeID:     ub.BadgeID(),
		Description: description,
		Level:       &level,
		EarnedAt:    ub.EarnedAt(),
		CreatedAt:   ub.CreatedAt(),
		UpdatedAt:   ub.UpdatedAt(),
	}
}

// UserBadgeModelToDomain 将 Model UserBadge 转换为 Domain UserBadge
func UserBadgeModelToDomain(m *models.UserBadge) *user_badges.UserBadge {
	if m == nil {
		return nil
	}

	description := ""
	if m.Description != nil {
		description = *m.Description
	}

	level := 1
	if m.Level != nil {
		level = *m.Level
	}

	state := user_badges.UserBadgeState{
		ID:          m.ID,
		UserID:      m.UserID,
		BadgeID:     m.BadgeID,
		Description: description,
		Level:       level,
		EarnedAt:    m.EarnedAt,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}

	return user_badges.NewUserBadgeFromState(state)
}

// UserBadgeModelToUpdates 将 Model UserBadge 转换为更新字段映射
func UserBadgeModelToUpdates(m *models.UserBadge) map[string]any {
	return map[string]any{
		models.UserBadgeCols.UserID:      m.UserID,
		models.UserBadgeCols.BadgeID:     m.BadgeID,
		models.UserBadgeCols.Description: m.Description,
		models.UserBadgeCols.Level:       m.Level,
		models.UserBadgeCols.EarnedAt:    m.EarnedAt,
		models.UserBadgeCols.UpdatedAt:   m.UpdatedAt,
	}
}
