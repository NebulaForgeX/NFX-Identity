package get

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/user_credentials"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/user_credentials/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 UserCredential，实现 user_credentials.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*user_credentials.UserCredential, error) {
	var m models.UserCredential
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_credentials.ErrUserCredentialNotFound
		}
		return nil, err
	}
	return mapper.UserCredentialModelToDomain(&m), nil
}
