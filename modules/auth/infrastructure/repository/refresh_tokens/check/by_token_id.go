package check

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// ByTokenID 根据 TokenID 检查 RefreshToken 是否存在，实现 refresh_tokens.Check 接口
func (h *Handler) ByTokenID(ctx context.Context, tokenID string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.RefreshToken{}).
		Where("token_id = ?", tokenID).
		Count(&count).Error
	return count > 0, err
}
