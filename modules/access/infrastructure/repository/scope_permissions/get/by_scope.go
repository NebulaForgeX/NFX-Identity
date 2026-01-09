package get

import (
	"context"
	"nfxid/modules/access/domain/scope_permissions"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/scope_permissions/mapper"
)

// ByScope 根据 Scope 获取 ScopePermission 列表，实现 scope_permissions.Get 接口
func (h *Handler) ByScope(ctx context.Context, scope string) ([]*scope_permissions.ScopePermission, error) {
	var ms []models.ScopePermission
	if err := h.db.WithContext(ctx).
		Where("scope = ?", scope).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*scope_permissions.ScopePermission, len(ms))
	for i, m := range ms {
		result[i] = mapper.ScopePermissionModelToDomain(&m)
	}
	return result, nil
}
