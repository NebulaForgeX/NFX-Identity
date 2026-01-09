package mapper

import (
	"nfxid/modules/image/domain/image_variants"
	"nfxid/modules/image/infrastructure/rdb/models"
)

// ImageVariantDomainToModel 将 Domain ImageVariant 转换为 Model ImageVariant
func ImageVariantDomainToModel(iv *image_variants.ImageVariant) *models.ImageVariant {
	if iv == nil {
		return nil
	}

	return &models.ImageVariant{
		ID:          iv.ID(),
		ImageID:     iv.ImageID(),
		VariantKey:  iv.VariantKey(),
		Width:       iv.Width(),
		Height:      iv.Height(),
		Size:        iv.Size(),
		MimeType:    iv.MimeType(),
		StoragePath: iv.StoragePath(),
		URL:         iv.URL(),
		CreatedAt:   iv.CreatedAt(),
		UpdatedAt:   iv.UpdatedAt(),
	}
}

// ImageVariantModelToDomain 将 Model ImageVariant 转换为 Domain ImageVariant
func ImageVariantModelToDomain(m *models.ImageVariant) *image_variants.ImageVariant {
	if m == nil {
		return nil
	}

	state := image_variants.ImageVariantState{
		ID:          m.ID,
		ImageID:     m.ImageID,
		VariantKey:  m.VariantKey,
		Width:       m.Width,
		Height:      m.Height,
		Size:        m.Size,
		MimeType:    m.MimeType,
		StoragePath: m.StoragePath,
		URL:         m.URL,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}

	return image_variants.NewImageVariantFromState(state)
}

// ImageVariantModelToUpdates 将 Model ImageVariant 转换为更新字段映射
func ImageVariantModelToUpdates(m *models.ImageVariant) map[string]any {
	return map[string]any{
		models.ImageVariantCols.ImageID:     m.ImageID,
		models.ImageVariantCols.VariantKey:  m.VariantKey,
		models.ImageVariantCols.Width:       m.Width,
		models.ImageVariantCols.Height:      m.Height,
		models.ImageVariantCols.Size:        m.Size,
		models.ImageVariantCols.MimeType:    m.MimeType,
		models.ImageVariantCols.StoragePath: m.StoragePath,
		models.ImageVariantCols.URL:         m.URL,
		models.ImageVariantCols.UpdatedAt:   m.UpdatedAt,
	}
}
