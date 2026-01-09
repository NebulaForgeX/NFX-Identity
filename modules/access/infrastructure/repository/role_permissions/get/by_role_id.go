package get

import (
	"context"
	"nfxid/modules/access/domain/role_permissions"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/role_permissions/mapper"

	"github.com/google/uuid"
)

// ByRoleID 根据 RoleID 获取 RolePermission 列表，实现 role_permissions.Get 接口
func (h *Handler) ByRoleID(ctx context.Context, roleID uuid.UUID) ([]*role_permissions.RolePermission, error) {
	var ms []models.RolePermission
	if err := h.db.WithContext(ctx).
		Where("role_id = ?", roleID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*role_permissions.RolePermission, len(ms))
	for i, m := range ms {
		result[i] = mapper.RolePermissionModelToDomain(&m)
	}
	return result, nil
}
