package reqdto

import (
	badgeAppCommands "nfxid/modules/directory/application/badges/commands"

	"github.com/google/uuid"
)

type BadgeCreateRequestDTO struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	IconURL     *string `json:"icon_url,omitempty"`
	Color       *string `json:"color,omitempty"`
	Category    *string `json:"category,omitempty"`
	IsSystem    bool    `json:"is_system,omitempty"`
}

type BadgeUpdateRequestDTO struct {
	ID          uuid.UUID `params:"id" validate:"required,uuid"`
	Name        string    `json:"name" validate:"required"`
	Description *string   `json:"description,omitempty"`
	IconURL     *string   `json:"icon_url,omitempty"`
	Color       *string   `json:"color,omitempty"`
	Category    *string   `json:"category,omitempty"`
}

type BadgeByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type BadgeByNameRequestDTO struct {
	Name string `params:"name" validate:"required"`
}

func (r *BadgeCreateRequestDTO) ToCreateCmd() badgeAppCommands.CreateBadgeCmd {
	return badgeAppCommands.CreateBadgeCmd{
		Name:        r.Name,
		Description: r.Description,
		IconURL:     r.IconURL,
		Color:       r.Color,
		Category:    r.Category,
		IsSystem:    r.IsSystem,
	}
}

func (r *BadgeUpdateRequestDTO) ToUpdateCmd() badgeAppCommands.UpdateBadgeCmd {
	return badgeAppCommands.UpdateBadgeCmd{
		BadgeID:     r.ID,
		Name:        r.Name,
		Description: r.Description,
		IconURL:     r.IconURL,
		Color:       r.Color,
		Category:    r.Category,
	}
}
