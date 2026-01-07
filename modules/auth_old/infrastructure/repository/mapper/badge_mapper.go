package mapper

import (
	"nfxid/modules/auth/domain/badge"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

func BadgeDomainToModel(b *badge.Badge) *models.Badge {
	return &models.Badge{
		ID:          b.ID(),
		Name:        b.Editable().Name,
		Description: b.Editable().Description,
		IconURL:     b.Editable().IconURL,
		Color:       b.Editable().Color,
		Category:    b.Editable().Category,
		IsSystem:    b.Editable().IsSystem,
		DeletedAt:   timex.TimeToGormDeletedAt(b.DeletedAt()),
	}
}

func BadgeModelToDomain(m *models.Badge) *badge.Badge {
	return badge.NewBadgeFromState(badge.BadgeState{
		ID: m.ID,
		Editable: badge.BadgeEditable{
			Name:        m.Name,
			Description: m.Description,
			IconURL:     m.IconURL,
			Color:       m.Color,
			Category:    m.Category,
			IsSystem:    m.IsSystem,
		},
		DeletedAt: timex.GormDeletedAtToTime(m.DeletedAt),
	})
}

func BadgeModelsToUpdates(m *models.Badge) map[string]any {
	return map[string]any{
		models.BadgeCols.Name:        m.Name,
		models.BadgeCols.Description: m.Description,
		models.BadgeCols.IconURL:     m.IconURL,
		models.BadgeCols.Color:       m.Color,
		models.BadgeCols.Category:    m.Category,
		models.BadgeCols.IsSystem:    m.IsSystem,
		models.BadgeCols.DeletedAt:   m.DeletedAt,
	}
}
