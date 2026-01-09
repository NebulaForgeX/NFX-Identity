package check

import (
	"context"
	"nfxid/modules/clients/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByAppIDAndScope 根据 AppID 和 Scope 检查 ClientScope 是否存在，实现 client_scopes.Check 接口
func (h *Handler) ByAppIDAndScope(ctx context.Context, appID uuid.UUID, scope string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.ClientScope{}).
		Where("app_id = ? AND scope = ?", appID, scope).
		Count(&count).Error
	return count > 0, err
}
