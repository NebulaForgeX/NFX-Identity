package get

import (
	"context"
	"nfxid/modules/access/domain/scope_permissions"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/scope_permissions/mapper"

	"github.com/google/uuid"
)

// ByPermissionID 根据 PermissionID 获取 ScopePermission 列表，实现 scope_permissions.Get 接口
func (h *Handler) ByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]*scope_permissions.ScopePermission, error) {
	var ms []models.ScopePermission
	if err := h.db.WithContext(ctx).
		Where("permission_id = ?", permissionID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*scope_permissions.ScopePermission, len(ms))
	for i, m := range ms {
		result[i] = mapper.ScopePermissionModelToDomain(&m)
	}
	return result, nil
}
