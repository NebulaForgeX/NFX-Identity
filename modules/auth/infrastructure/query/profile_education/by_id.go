package profile_education

import (
	"context"
	"errors"
	educationDomainErrors "nfxid/modules/auth/domain/profile_education/errors"
	educationDomainViews "nfxid/modules/auth/domain/profile_education/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Education，实现 education.Query 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (educationDomainViews.EducationView, error) {
	var m models.Education
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return educationDomainViews.EducationView{}, educationDomainErrors.ErrEducationViewNotFound
		}
		return educationDomainViews.EducationView{}, err
	}
	return mapper.EducationModelToDomain(&m), nil
}
