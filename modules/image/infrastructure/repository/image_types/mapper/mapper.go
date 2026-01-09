package mapper

import (
	"nfxid/modules/image/domain/image_types"
	"nfxid/modules/image/infrastructure/rdb/models"
)

// ImageTypeDomainToModel 将 Domain ImageType 转换为 Model ImageType
func ImageTypeDomainToModel(it *image_types.ImageType) *models.ImageType {
	if it == nil {
		return nil
	}

	isSystem := it.IsSystem()
	return &models.ImageType{
		ID:          it.ID(),
		Key:         it.Key(),
		Description: it.Description(),
		MaxWidth:    it.MaxWidth(),
		MaxHeight:   it.MaxHeight(),
		AspectRatio: it.AspectRatio(),
		IsSystem:    &isSystem,
		CreatedAt:   it.CreatedAt(),
		UpdatedAt:   it.UpdatedAt(),
	}
}

// ImageTypeModelToDomain 将 Model ImageType 转换为 Domain ImageType
func ImageTypeModelToDomain(m *models.ImageType) *image_types.ImageType {
	if m == nil {
		return nil
	}

	isSystem := false
	if m.IsSystem != nil {
		isSystem = *m.IsSystem
	}

	state := image_types.ImageTypeState{
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

	return image_types.NewImageTypeFromState(state)
}

// ImageTypeModelToUpdates 将 Model ImageType 转换为更新字段映射
func ImageTypeModelToUpdates(m *models.ImageType) map[string]any {
	return map[string]any{
		models.ImageTypeCols.Key:         m.Key,
		models.ImageTypeCols.Description: m.Description,
		models.ImageTypeCols.MaxWidth:    m.MaxWidth,
		models.ImageTypeCols.MaxHeight:   m.MaxHeight,
		models.ImageTypeCols.AspectRatio: m.AspectRatio,
		models.ImageTypeCols.IsSystem:    m.IsSystem,
		models.ImageTypeCols.UpdatedAt:   m.UpdatedAt,
	}
}
