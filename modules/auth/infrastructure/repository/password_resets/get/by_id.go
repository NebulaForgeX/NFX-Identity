package get

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/password_resets"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/password_resets/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 PasswordReset，实现 password_resets.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*password_resets.PasswordReset, error) {
	var m models.PasswordReset
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, password_resets.ErrPasswordResetNotFound
		}
		return nil, err
	}
	return mapper.PasswordResetModelToDomain(&m), nil
}
