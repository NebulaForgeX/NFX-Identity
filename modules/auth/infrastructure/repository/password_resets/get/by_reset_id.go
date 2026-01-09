package get

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/password_resets"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/password_resets/mapper"

	"gorm.io/gorm"
)

// ByResetID 根据 ResetID 获取 PasswordReset，实现 password_resets.Get 接口
func (h *Handler) ByResetID(ctx context.Context, resetID string) (*password_resets.PasswordReset, error) {
	var m models.PasswordReset
	if err := h.db.WithContext(ctx).Where("reset_id = ?", resetID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, password_resets.ErrPasswordResetNotFound
		}
		return nil, err
	}
	return mapper.PasswordResetModelToDomain(&m), nil
}
