package update

import (
	"context"
	"nfxid/modules/auth/domain/refresh_tokens"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/refresh_tokens/mapper"
)

// Generic 通用更新 RefreshToken，实现 refresh_tokens.Update 接口
func (h *Handler) Generic(ctx context.Context, rt *refresh_tokens.RefreshToken) error {
	m := mapper.RefreshTokenDomainToModel(rt)
	updates := mapper.RefreshTokenModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.RefreshToken{}).
		Where("id = ?", rt.ID()).
		Updates(updates).Error
}
