package delete

import (
	"context"
	"nfxid/modules/access/domain/scopes"
	"nfxid/modules/access/infrastructure/rdb/models"
)

// ByScope 根据 Scope 删除 Scope，实现 scopes.Delete 接口
func (h *Handler) ByScope(ctx context.Context, scope string) error {
	result := h.db.WithContext(ctx).
		Where("scope = ?", scope).
		Delete(&models.Scope{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return scopes.ErrScopeNotFound
	}
	return nil
}
