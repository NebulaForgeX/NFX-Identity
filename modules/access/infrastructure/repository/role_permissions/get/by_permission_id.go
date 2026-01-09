package get

import (
	"context"
	"nfxid/modules/access/domain/role_permissions"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/role_permissions/mapper"

	"github.com/google/uuid"
)

// ByPermissionID 根据 PermissionID 获取 RolePermission 列表，实现 role_permissions.Get 接口
func (h *Handler) ByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]*role_permissions.RolePermission, error) {
	var ms []models.RolePermission
	if err := h.db.WithContext(ctx).
		Where("permission_id = ?", permissionID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*role_permissions.RolePermission, len(ms))
	for i, m := range ms {
		result[i] = mapper.RolePermissionModelToDomain(&m)
	}
	return result, nil
}
