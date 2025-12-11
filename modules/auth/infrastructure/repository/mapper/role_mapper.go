package mapper

import (
	"encoding/json"
	"nfxid/modules/auth/domain/role"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"

	"gorm.io/datatypes"
)

func RoleDomainToModel(r *role.Role) *models.Role {
	if r == nil {
		return nil
	}

	editable := r.Editable()

	// 序列化 Permissions JSON
	var permissions *datatypes.JSON
	if len(editable.Permissions) > 0 {
		data, _ := json.Marshal(editable.Permissions)
		jsonData := datatypes.JSON(data)
		permissions = &jsonData
	}

	return &models.Role{
		ID:          r.ID(),
		Name:        editable.Name,
		Description: editable.Description,
		Permissions: permissions,
		IsSystem:    r.IsSystem(),
		CreatedAt:   r.CreatedAt(),
		UpdatedAt:   r.UpdatedAt(),
		DeletedAt:   timex.TimeToGormDeletedAt(r.DeletedAt()),
	}
}

func RoleModelToDomain(m *models.Role) *role.Role {
	if m == nil {
		return nil
	}

	// 解析 Permissions JSON
	var permissions []string
	if m.Permissions != nil {
		json.Unmarshal(*m.Permissions, &permissions)
	}

	editable := role.RoleEditable{
		Name:        m.Name,
		Description: m.Description,
		Permissions: permissions,
	}

	state := role.RoleState{
		ID:        m.ID,
		Editable:  editable,
		IsSystem:  m.IsSystem,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return role.NewRoleFromState(state)
}

func RoleModelsToUpdates(m *models.Role) map[string]any {
	return map[string]any{
		models.RoleCols.Name:        m.Name,
		models.RoleCols.Description: m.Description,
		models.RoleCols.Permissions: m.Permissions,
		models.RoleCols.IsSystem:    m.IsSystem,
		models.RoleCols.DeletedAt:   m.DeletedAt,
	}
}
