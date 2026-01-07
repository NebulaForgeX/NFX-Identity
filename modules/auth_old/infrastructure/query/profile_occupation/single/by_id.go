package single

import (
	"context"
	"errors"
	occupationDomainErrors "nfxid/modules/auth/domain/profile_occupation/errors"
	occupationDomainViews "nfxid/modules/auth/domain/profile_occupation/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Occupation，实现 occupationDomain.Single 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*occupationDomainViews.OccupationView, error) {
	var m models.Occupation
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, occupationDomainErrors.ErrOccupationViewNotFound
		}
		return nil, err
	}
	view := mapper.OccupationModelToDomain(&m)
	return &view, nil
}
