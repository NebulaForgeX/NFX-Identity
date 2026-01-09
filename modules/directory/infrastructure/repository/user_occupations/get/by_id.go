package get

import (
	"context"
	"errors"
	"nfxid/modules/directory/domain/user_occupations"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_occupations/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 UserOccupation，实现 user_occupations.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*user_occupations.UserOccupation, error) {
	var m models.UserOccupation
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_occupations.ErrUserOccupationNotFound
		}
		return nil, err
	}
	return mapper.UserOccupationModelToDomain(&m), nil
}
