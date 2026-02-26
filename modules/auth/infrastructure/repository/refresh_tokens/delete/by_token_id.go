package delete

import (
	"context"
	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
)

// ByTokenID 根据 TokenID 删除 RefreshToken，实现 refresh_tokens.Delete 接口
func (h *Handler) ByTokenID(ctx context.Context, tokenID string) error {
	result := h.db.WithContext(ctx).
		Where("token_id = ?", tokenID).
		Delete(&models.RefreshToken{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return authErr.ErrRefreshTokenNotFound
	}
	return nil
}
