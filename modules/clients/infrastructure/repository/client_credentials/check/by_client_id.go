package check

import (
	"context"
	"nfxid/modules/clients/infrastructure/rdb/models"
)

// ByClientID 根据 ClientID 检查 ClientCredential 是否存在，实现 client_credentials.Check 接口
func (h *Handler) ByClientID(ctx context.Context, clientID string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.ClientCredential{}).
		Where("client_id = ?", clientID).
		Count(&count).Error
	return count > 0, err
}
