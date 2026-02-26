package update

import (
	"context"
	"errors"
	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/domain/refresh_tokens"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/refresh_tokens/mapper"
	"time"

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
			return authErr.ErrRefreshTokenNotFound
		}
		return err
	}

	// 检查是否已经撤销
	if m.RevokedAt != nil {
		return authErr.ErrTokenAlreadyRevoked
	}

	now := time.Now().UTC()
	revokeReason := mapper.RevokeReasonDomainToEnum(reason)
	updates := map[string]any{
		models.RefreshTokenCols.RevokedAt:    &now,
		models.RefreshTokenCols.RevokeReason: &revokeReason,
		models.RefreshTokenCols.UpdatedAt:    now,
	}

	return h.db.WithContext(ctx).
		Model(&models.RefreshToken{}).
		Where("token_id = ?", tokenID).
		Updates(updates).Error
}
