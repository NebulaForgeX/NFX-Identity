package views

import (
	"time"

	badgeDomainViews "nebulaid/modules/auth/domain/badge/views"

	"github.com/google/uuid"
)

type BadgeView struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	IconURL     *string   `json:"icon_url"`
	Color       *string   `json:"color"`
	Category    *string   `json:"category"`
	IsSystem    bool      `json:"is_system"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// BadgeViewMapper 将 Domain BadgeView 转换为 Application BadgeView
func BadgeViewMapper(v badgeDomainViews.BadgeView) BadgeView {
	return BadgeView{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
		IconURL:     v.IconURL,
		Color:       v.Color,
		Category:    v.Category,
		IsSystem:    v.IsSystem,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}
}
