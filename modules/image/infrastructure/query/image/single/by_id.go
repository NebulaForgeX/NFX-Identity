package single

import (
	"context"
	"encoding/json"
	"errors"
	imageDomainErrors "nfxid/modules/image/domain/image/errors"
	imageDomainViews "nfxid/modules/image/domain/image/views"
	"nfxid/modules/image/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Image，实现 imageDomain.Single 接口
func (h *Handler) ByID(ctx context.Context, imageID uuid.UUID) (*imageDomainViews.ImageView, error) {
	var m models.Image
	if err := h.db.WithContext(ctx).Where("id = ?", imageID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageDomainErrors.ErrImageNotFound
		}
		return nil, err
	}
	view := imageModelToDomainView(&m)
	return &view, nil
}

func imageModelToDomainView(m *models.Image) imageDomainViews.ImageView {
	var metadata map[string]interface{}
	if m.Metadata != nil {
		rawJSON, _ := m.Metadata.MarshalJSON()
		if len(rawJSON) > 0 {
			_ = json.Unmarshal(rawJSON, &metadata)
		}
	}

	return imageDomainViews.ImageView{
		ID:               m.ID,
		TypeID:           m.TypeID,
		UserID:           m.UserID,
		SourceDomain:     m.SourceDomain,
		Filename:         m.Filename,
		OriginalFilename: m.OriginalFilename,
		MimeType:         m.MimeType,
		Size:             m.Size,
		Width:            m.Width,
		Height:           m.Height,
		StoragePath:      m.StoragePath,
		URL:              m.URL,
		IsPublic:         m.IsPublic,
		Metadata:         metadata,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
	}
}
