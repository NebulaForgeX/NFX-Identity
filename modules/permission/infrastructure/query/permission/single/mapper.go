package single

import (
	permissionDomainViews "nfxid/modules/permission/domain/permission/views"
	"nfxid/modules/permission/infrastructure/rdb/models"
)

func permissionModelToDomainView(m *models.Permission) permissionDomainViews.PermissionView {
	description := ""
	if m.Description != nil {
		description = *m.Description
	}

	return permissionDomainViews.PermissionView{
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
