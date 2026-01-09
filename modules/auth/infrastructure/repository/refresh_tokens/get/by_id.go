package get

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/refresh_tokens"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/refresh_tokens/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 RefreshToken，实现 refresh_tokens.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*refresh_tokens.RefreshToken, error) {
	var m models.RefreshToken
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, refresh_tokens.ErrRefreshTokenNotFound
		}
		return nil, err
	}
	return mapper.RefreshTokenModelToDomain(&m), nil
}
