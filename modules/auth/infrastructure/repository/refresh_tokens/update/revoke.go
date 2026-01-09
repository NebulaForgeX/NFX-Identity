package update

import (
	"context"
	"errors"
	"time"
	"nfxid/modules/auth/domain/refresh_tokens"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/refresh_tokens/mapper"

	"gorm.io/gorm"
)

// Revoke 撤销 RefreshToken，实现 refresh_tokens.Update 接口
func (h *Handler) Revoke(ctx context.Context, tokenID string, reason refresh_tokens.RevokeReason) error {
	// 先检查 RefreshToken 是否存在
	var m models.RefreshToken
	if err := h.db.WithContext(ctx).
		Where("token_id = ?", tokenID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return refresh_tokens.ErrRefreshTokenNotFound
		}
		return err
	}

	// 检查是否已经撤销
	if m.RevokedAt != nil {
		return refresh_tokens.ErrTokenAlreadyRevoked
	}

	now := time.Now().UTC()
	revokeReason := mapper.RevokeReasonDomainToEnum(reason)
	updates := map[string]any{
		models.RefreshTokenCols.RevokedAt:    &now,
		models.RefreshTokenCols.RevokeReason: &revokeReason,
		models.RefreshTokenCols.UpdatedAt:   now,
	}

	return h.db.WithContext(ctx).
		Model(&models.RefreshToken{}).
		Where("token_id = ?", tokenID).
		Updates(updates).Error
}
