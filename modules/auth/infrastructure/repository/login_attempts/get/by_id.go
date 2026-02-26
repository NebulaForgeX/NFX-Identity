package get

import (
	"context"
	"errors"
	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/domain/login_attempts"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/login_attempts/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 LoginAttempt，实现 login_attempts.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*login_attempts.LoginAttempt, error) {
	var m models.LoginAttempt
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrLoginAttemptNotFound
		}
		return nil, err
	}
	return mapper.LoginAttemptModelToDomain(&m), nil
}
