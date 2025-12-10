package reqdto

import (
	roleAppCommands "nebulaid/modules/auth/application/role/commands"
	roleAppQueries "nebulaid/modules/auth/application/role/queries"
	roleDomain "nebulaid/modules/auth/domain/role"
	"nebulaid/pkgs/query"
	"nebulaid/pkgs/utils/ptr"

	"github.com/google/uuid"
)

type RoleCreateRequestDTO struct {
	Name        string   `json:"name" validate:"required"`
	Description *string  `json:"description,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	IsSystem    bool     `json:"is_system"`
}

type RoleUpdateRequestDTO struct {
	ID          uuid.UUID `params:"id" validate:"required,uuid"`
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Permissions *[]string `json:"permissions,omitempty"`
	IsSystem    *bool     `json:"is_system,omitempty"`
}

type RoleByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type RoleByNameRequestDTO struct {
	Name string `params:"name" validate:"required"`
}

type RoleQueryParamsDTO struct {
	Offset   *int     `query:"offset"`
	Limit    *int     `query:"limit"`
	Search   *string  `query:"search"`
	IsSystem *bool    `query:"is_system"`
	Sort     []string `query:"sort"`
}

func (r *RoleCreateRequestDTO) ToCreateCmd() roleAppCommands.CreateRoleCmd {
	return roleAppCommands.CreateRoleCmd{
		Editable: roleDomain.RoleEditable{
			Name:        r.Name,
			Description: r.Description,
			Permissions: r.Permissions,
		},
		IsSystem: r.IsSystem,
	}
}

func (r *RoleUpdateRequestDTO) ToUpdateCmd() roleAppCommands.UpdateRoleCmd {
	editable := roleDomain.RoleEditable{}
	if r.Name != nil {
		editable.Name = *r.Name
	}
	if r.Description != nil {
		editable.Description = r.Description
	}
	if r.Permissions != nil {
		editable.Permissions = *r.Permissions
	}

	return roleAppCommands.UpdateRoleCmd{
		RoleID:   r.ID,
		Editable: editable,
	}
}

func (r *RoleQueryParamsDTO) ToListQuery() roleAppQueries.RoleListQuery {
	return roleAppQueries.RoleListQuery{
		DomainPagination: query.DomainPagination{
			Offset: ptr.Deref(r.Offset),
			Limit:  ptr.Deref(r.Limit),
		},
		DomainSorts: query.ParseSortParams(r.Sort, map[string]roleAppQueries.SortField{
			"created_time": roleAppQueries.SortByCreatedTime,
			"name":         roleAppQueries.SortByName,
		}),
		Search:   r.Search,
		IsSystem: r.IsSystem,
	}
}
