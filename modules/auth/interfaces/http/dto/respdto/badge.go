package respdto

import (
	"time"

	badgeAppViews "nfxid/modules/auth/application/badge/views"

	"github.com/google/uuid"
)

type BadgeDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	IconURL     *string   `json:"icon_url,omitempty"`
	Color       *string   `json:"color,omitempty"`
	Category    *string   `json:"category,omitempty"`
	IsSystem    bool      `json:"is_system"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// BadgeViewToDTO converts application BadgeView to response DTO
func BadgeViewToDTO(v *badgeAppViews.BadgeView) *BadgeDTO {
	if v == nil {
		return nil
	}

	return &BadgeDTO{
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

// BadgeListViewToDTO converts list of BadgeView to DTOs
func BadgeListViewToDTO(views []badgeAppViews.BadgeView) []BadgeDTO {
	dtos := make([]BadgeDTO, len(views))
	for i, v := range views {
		if dto := BadgeViewToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
