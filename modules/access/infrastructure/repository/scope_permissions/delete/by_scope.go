package delete

import (
	"context"
	"nfxid/modules/access/infrastructure/rdb/models"
)

// ByScope 根据 Scope 删除所有 ScopePermission，实现 scope_permissions.Delete 接口
func (h *Handler) ByScope(ctx context.Context, scope string) error {
	return h.db.WithContext(ctx).
		Where("scope = ?", scope).
		Delete(&models.ScopePermission{}).Error
}
