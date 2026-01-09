package get

import (
	"context"
	"errors"
	"nfxid/modules/directory/domain/user_educations"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_educations/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 UserEducation，实现 user_educations.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*user_educations.UserEducation, error) {
	var m models.UserEducation
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_educations.ErrUserEducationNotFound
		}
		return nil, err
	}
	return mapper.UserEducationModelToDomain(&m), nil
}
