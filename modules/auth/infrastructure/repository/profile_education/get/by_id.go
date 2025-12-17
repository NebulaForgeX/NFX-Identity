package get

import (
	"context"
	"errors"
	education "nfxid/modules/auth/domain/profile_education"
	educationDomainErrors "nfxid/modules/auth/domain/profile_education/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Education，实现 education.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*education.Education, error) {
	var m models.Education
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, educationDomainErrors.ErrEducationNotFound
		}
		return nil, err
	}
	return mapper.EducationModelToDomain(&m), nil
}
