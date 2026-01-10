package respdto

import (
	"time"

	badgeAppResult "nfxid/modules/directory/application/badges/results"

	"github.com/google/uuid"
)

type BadgeDTO struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	IconURL     *string    `json:"icon_url,omitempty"`
	Color       *string    `json:"color,omitempty"`
	Category    *string    `json:"category,omitempty"`
	IsSystem    bool       `json:"is_system"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// BadgeROToDTO converts application BadgeRO to response DTO
func BadgeROToDTO(v *badgeAppResult.BadgeRO) *BadgeDTO {
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
		DeletedAt:   v.DeletedAt,
	}
}

// BadgeListROToDTO converts list of BadgeRO to DTOs
func BadgeListROToDTO(results []badgeAppResult.BadgeRO) []BadgeDTO {
	dtos := make([]BadgeDTO, len(results))
	for i, v := range results {
		if dto := BadgeROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
