package create

import (
	"context"
	"nfxid/modules/auth/domain/refresh_tokens"
	"nfxid/modules/auth/infrastructure/repository/refresh_tokens/mapper"
)

// New 创建新的 RefreshToken，实现 refresh_tokens.Create 接口
func (h *Handler) New(ctx context.Context, rt *refresh_tokens.RefreshToken) error {
	m := mapper.RefreshTokenDomainToModel(rt)
	return h.db.WithContext(ctx).Create(&m).Error
}
