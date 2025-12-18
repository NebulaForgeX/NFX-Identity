package list

import (
	permissionDomainViews "nfxid/modules/permission/domain/permission/views"
	"nfxid/modules/permission/infrastructure/rdb/models"
)

func permissionModelToDomainView(m *models.Permission) permissionDomainViews.PermissionView {
	description := ""
	if m.Description != nil {
		description = *m.Description
	}
	category := ""
	if m.Category != nil {
		category = *m.Category
	}

	return permissionDomainViews.PermissionView{
		ID:          m.ID,
		Tag:         m.Tag,
		Name:        m.Name,
		Description: description,
		Category:    category,
		IsSystem:    m.IsSystem,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
