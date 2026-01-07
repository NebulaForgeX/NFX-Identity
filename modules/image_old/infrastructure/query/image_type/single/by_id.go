package single

import (
	"context"
	"errors"
	imageTypeDomainErrors "nfxid/modules/image/domain/image_type/errors"
	imageTypeDomainViews "nfxid/modules/image/domain/image_type/views"
	"nfxid/modules/image/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 ImageType，实现 imageTypeDomain.Single 接口
func (h *Handler) ByID(ctx context.Context, imageTypeID uuid.UUID) (*imageTypeDomainViews.ImageTypeView, error) {
	var m models.ImageType
	if err := h.db.WithContext(ctx).Where("id = ?", imageTypeID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageTypeDomainErrors.ErrImageTypeNotFound
		}
		return nil, err
	}
	view := imageTypeModelToDomainView(&m)
	return &view, nil
}

// ByKey 根据 Key 获取 ImageType，实现 imageTypeDomain.Single 接口
func (h *Handler) ByKey(ctx context.Context, key string) (*imageTypeDomainViews.ImageTypeView, error) {
	var m models.ImageType
	if err := h.db.WithContext(ctx).Where("key = ?", key).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageTypeDomainErrors.ErrImageTypeNotFound
		}
		return nil, err
	}
	view := imageTypeModelToDomainView(&m)
	return &view, nil
}

func imageTypeModelToDomainView(m *models.ImageType) imageTypeDomainViews.ImageTypeView {
	isSystem := false
	if m.IsSystem != nil {
		isSystem = *m.IsSystem
	}

	return imageTypeDomainViews.ImageTypeView{
		ID:          m.ID,
		Key:         m.Key,
		Description: m.Description,
		MaxWidth:    m.MaxWidth,
		MaxHeight:   m.MaxHeight,
		AspectRatio: m.AspectRatio,
		IsSystem:    isSystem,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
