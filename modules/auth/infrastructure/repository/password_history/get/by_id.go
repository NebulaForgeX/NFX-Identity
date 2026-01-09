package get

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/password_history"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/password_history/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 PasswordHistory，实现 password_history.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*password_history.PasswordHistory, error) {
	var m models.PasswordHistory
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, password_history.ErrPasswordHistoryNotFound
		}
		return nil, err
	}
	return mapper.PasswordHistoryModelToDomain(&m), nil
}
