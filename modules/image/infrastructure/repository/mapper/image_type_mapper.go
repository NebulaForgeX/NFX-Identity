package mapper

import (
	imageTypeDomain "nebulaid/modules/image/domain/image_type"
	"nebulaid/modules/image/infrastructure/rdb/models"
)

func ImageTypeDomainToModel(it *imageTypeDomain.ImageType) *models.ImageType {
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

func ImageTypeModelToDomain(m *models.ImageType) *imageTypeDomain.ImageType {
	if m == nil {
		return nil
	}

	isSystem := false
	if m.IsSystem != nil {
		isSystem = *m.IsSystem
	}

	editable := imageTypeDomain.ImageTypeEditable{
		Key:         m.Key,
		Description: m.Description,
		MaxWidth:    m.MaxWidth,
		MaxHeight:   m.MaxHeight,
		AspectRatio: m.AspectRatio,
		IsSystem:    isSystem,
	}

	state := imageTypeDomain.ImageTypeState{
		ID:          m.ID,
		Key:         editable.Key,
		Description: editable.Description,
		MaxWidth:    editable.MaxWidth,
		MaxHeight:   editable.MaxHeight,
		AspectRatio: editable.AspectRatio,
		IsSystem:    editable.IsSystem,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}

	return imageTypeDomain.NewImageTypeFromState(state)
}

func ImageTypeModelsToUpdates(m *models.ImageType) map[string]any {
	return map[string]any{
		models.ImageTypeCols.Key:         m.Key,
		models.ImageTypeCols.Description: m.Description,
		models.ImageTypeCols.MaxWidth:    m.MaxWidth,
		models.ImageTypeCols.MaxHeight:   m.MaxHeight,
		models.ImageTypeCols.AspectRatio: m.AspectRatio,
		models.ImageTypeCols.IsSystem:    m.IsSystem,
	}
}
