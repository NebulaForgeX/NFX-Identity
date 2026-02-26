package update

import (
	"context"
	"nfxidentity/modules/auth/domain/refresh_tokens"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/refresh_tokens/mapper"
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
