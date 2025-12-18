package mapper

import (
	"nfxid/modules/permission/domain/permission"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

func PermissionDomainToModel(p *permission.Permission) *models.Permission {
	if p == nil {
		return nil
	}

	editable := p.Editable()
	description := editable.Description
	if description == "" {
		description = ""
	}
	var descriptionPtr *string
	if description != "" {
		descriptionPtr = &description
	}

	return &models.Permission{
		ID:          p.ID(),
		Tag:         editable.Tag,
		Name:        editable.Name,
		Description: descriptionPtr,
		Category:    editable.Category,
		IsSystem:    p.IsSystem(),
		CreatedAt:   p.CreatedAt(),
		UpdatedAt:   p.UpdatedAt(),
		DeletedAt:   timex.TimeToGormDeletedAt(p.DeletedAt()),
	}
}

func PermissionModelToDomain(m *models.Permission) *permission.Permission {
	if m == nil {
		return nil
	}

	description := ""
	if m.Description != nil {
		description = *m.Description
	}

	editable := permission.PermissionEditable{
		Tag:         m.Tag,
		Name:        m.Name,
		Description: description,
		Category:    m.Category,
	}

	state := permission.PermissionState{
		ID:        m.ID,
		Editable:  editable,
		IsSystem:  m.IsSystem,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return permission.NewPermissionFromState(state)
}

func PermissionModelsToUpdates(m *models.Permission) map[string]any {
	return map[string]any{
		models.PermissionCols.Tag:         m.Tag,
		models.PermissionCols.Name:        m.Name,
		models.PermissionCols.Description: m.Description,
		models.PermissionCols.Category:    m.Category,
		models.PermissionCols.IsSystem:    m.IsSystem,
		models.PermissionCols.DeletedAt:   m.DeletedAt,
	}
}
