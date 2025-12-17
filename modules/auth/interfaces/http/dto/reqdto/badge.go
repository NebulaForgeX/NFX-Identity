package reqdto

import (
	badgeAppCommands "nfxid/modules/auth/application/badge/commands"
	badgeDomain "nfxid/modules/auth/domain/badge"
	"nfxid/pkgs/query"
	"nfxid/pkgs/utils/ptr"

	"github.com/google/uuid"
)

type BadgeCreateRequestDTO struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	IconURL     *string `json:"icon_url,omitempty"`
	Color       *string `json:"color,omitempty"`
	Category    *string `json:"category,omitempty"`
	IsSystem    bool    `json:"is_system"`
}

type BadgeUpdateRequestDTO struct {
	ID          uuid.UUID `params:"id" validate:"required,uuid"`
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	IconURL     *string   `json:"icon_url,omitempty"`
	Color       *string   `json:"color,omitempty"`
	Category    *string   `json:"category,omitempty"`
	IsSystem    *bool     `json:"is_system,omitempty"`
}

type BadgeByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type BadgeByNameRequestDTO struct {
	Name string `params:"name" validate:"required"`
}

type BadgeQueryParamsDTO struct {
	Offset   *int     `query:"offset"`
	Limit    *int     `query:"limit"`
	Search   *string  `query:"search"`
	Category *string  `query:"category"`
	IsSystem *bool    `query:"is_system"`
	Sort     []string `query:"sort"`
}

func (r *BadgeCreateRequestDTO) ToCreateCmd() badgeAppCommands.CreateBadgeCmd {
	return badgeAppCommands.CreateBadgeCmd{
		Editable: badgeDomain.BadgeEditable{
			Name:        r.Name,
			Description: r.Description,
			IconURL:     r.IconURL,
			Color:       r.Color,
			Category:    r.Category,
			IsSystem:    r.IsSystem,
		},
	}
}

func (r *BadgeUpdateRequestDTO) ToUpdateCmd() badgeAppCommands.UpdateBadgeCmd {
	editable := badgeDomain.BadgeEditable{}
	if r.Name != nil {
		editable.Name = *r.Name
	}
	if r.Description != nil {
		editable.Description = r.Description
	}
	if r.IconURL != nil {
		editable.IconURL = r.IconURL
	}
	if r.Color != nil {
		editable.Color = r.Color
	}
	if r.Category != nil {
		editable.Category = r.Category
	}
	if r.IsSystem != nil {
		editable.IsSystem = *r.IsSystem
	}

	return badgeAppCommands.UpdateBadgeCmd{
		BadgeID:  r.ID,
		Editable: editable,
	}
}

func (r *BadgeQueryParamsDTO) ToListQuery() badgeDomain.ListQuery {
	return badgeDomain.ListQuery{
		DomainPagination: query.DomainPagination{
			Offset: ptr.Deref(r.Offset),
			Limit:  ptr.Deref(r.Limit),
		},
		DomainSorts: query.ParseSortParams(r.Sort, map[string]badgeDomain.SortField{
			"created_time": badgeDomain.SortByCreatedTime,
			"name":         badgeDomain.SortByName,
		}),
		Search:   r.Search,
		Category: r.Category,
		IsSystem: r.IsSystem,
	}
}
