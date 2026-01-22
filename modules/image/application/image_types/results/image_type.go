package results

import (
	"time"

	"nfxid/modules/image/domain/image_types"

	"github.com/google/uuid"
)

type ImageTypeRO struct {
	ID          uuid.UUID
	Key         string
	Description *string
	MaxWidth    *int
	MaxHeight   *int
	AspectRatio *string
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ImageTypeMapper 将 Domain ImageType 转换为 Application ImageTypeRO
func ImageTypeMapper(it *image_types.ImageType) ImageTypeRO {
	if it == nil {
		return ImageTypeRO{}
	}

	return ImageTypeRO{
		ID:          it.ID(),
		Key:         it.Key(),
		Description: it.Description(),
		MaxWidth:    it.MaxWidth(),
		MaxHeight:   it.MaxHeight(),
		AspectRatio: it.AspectRatio(),
		IsSystem:    it.IsSystem(),
		CreatedAt:   it.CreatedAt(),
		UpdatedAt:   it.UpdatedAt(),
	}
}
