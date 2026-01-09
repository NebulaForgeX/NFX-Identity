package mapper

import (
	"nfxid/modules/directory/domain/badges"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

// BadgeDomainToModel 将 Domain Badge 转换为 Model Badge
func BadgeDomainToModel(b *badges.Badge) *models.Badge {
	if b == nil {
		return nil
	}

	return &models.Badge{
		ID:          b.ID(),
		Name:        b.Name(),
		Description: b.Description(),
		IconURL:     b.IconURL(),
		Color:       b.Color(),
		Category:    b.Category(),
		IsSystem:    b.IsSystem(),
		CreatedAt:   b.CreatedAt(),
		UpdatedAt:   b.UpdatedAt(),
		DeletedAt:   timex.TimeToGormDeletedAt(b.DeletedAt()),
	}
}

// BadgeModelToDomain 将 Model Badge 转换为 Domain Badge
func BadgeModelToDomain(m *models.Badge) *badges.Badge {
	if m == nil {
		return nil
	}

	state := badges.BadgeState{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		IconURL:     m.IconURL,
		Color:       m.Color,
		Category:    m.Category,
		IsSystem:    m.IsSystem,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return badges.NewBadgeFromState(state)
}

// BadgeModelToUpdates 将 Model Badge 转换为更新字段映射
func BadgeModelToUpdates(m *models.Badge) map[string]any {
	return map[string]any{
		models.BadgeCols.Name:        m.Name,
		models.BadgeCols.Description: m.Description,
		models.BadgeCols.IconURL:     m.IconURL,
		models.BadgeCols.Color:       m.Color,
		models.BadgeCols.Category:    m.Category,
		models.BadgeCols.IsSystem:    m.IsSystem,
		models.BadgeCols.UpdatedAt:   m.UpdatedAt,
		models.BadgeCols.DeletedAt:   m.DeletedAt,
	}
}
