package get

import (
	"context"
	"errors"
	occupation "nfxid/modules/auth/domain/profile_occupation"
	occupationDomainErrors "nfxid/modules/auth/domain/profile_occupation/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Occupation，实现 occupation.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*occupation.Occupation, error) {
	var m models.Occupation
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, occupationDomainErrors.ErrOccupationNotFound
		}
		return nil, err
	}
	return mapper.OccupationModelToDomain(&m), nil
}
