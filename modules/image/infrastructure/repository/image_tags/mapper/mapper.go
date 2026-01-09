package mapper

import (
	"nfxid/modules/image/domain/image_tags"
	"nfxid/modules/image/infrastructure/rdb/models"
)

// ImageTagDomainToModel 将 Domain ImageTag 转换为 Model ImageTag
func ImageTagDomainToModel(it *image_tags.ImageTag) *models.ImageTag {
	if it == nil {
		return nil
	}

	return &models.ImageTag{
		ID:         it.ID(),
		ImageID:    it.ImageID(),
		Tag:        it.Tag(),
		Confidence: it.Confidence(),
		CreatedAt:  it.CreatedAt(),
	}
}

// ImageTagModelToDomain 将 Model ImageTag 转换为 Domain ImageTag
func ImageTagModelToDomain(m *models.ImageTag) *image_tags.ImageTag {
	if m == nil {
		return nil
	}

	state := image_tags.ImageTagState{
		ID:         m.ID,
		ImageID:    m.ImageID,
		Tag:        m.Tag,
		Confidence: m.Confidence,
		CreatedAt:  m.CreatedAt,
	}

	return image_tags.NewImageTagFromState(state)
}

// ImageTagModelToUpdates 将 Model ImageTag 转换为更新字段映射
func ImageTagModelToUpdates(m *models.ImageTag) map[string]any {
	return map[string]any{
		models.ImageTagCols.ImageID:    m.ImageID,
		models.ImageTagCols.Tag:        m.Tag,
		models.ImageTagCols.Confidence: m.Confidence,
	}
}
