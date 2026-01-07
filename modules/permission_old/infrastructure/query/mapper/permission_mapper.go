package mapper

import (
	permissionAppViews "nfxid/modules/permission/application/permission/views"
	"nfxid/modules/permission/infrastructure/rdb/models"
)

func PermissionModelToAppView(m *models.Permission) permissionAppViews.PermissionView {
	description := ""
	if m.Description != nil {
		description = *m.Description
	}

	return permissionAppViews.PermissionView{
		ID:          m.ID,
		Tag:         m.Tag,
		Name:        m.Name,
		Description: description,
		Category:    m.Category,
		IsSystem:    m.IsSystem,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
