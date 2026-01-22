package results

import (
	"time"

	"nfxid/modules/image/domain/image_tags"

	"github.com/google/uuid"
)

type ImageTagRO struct {
	ID         uuid.UUID
	ImageID    uuid.UUID
	Tag        string
	Confidence *float64
	CreatedAt  time.Time
}

// ImageTagMapper 将 Domain ImageTag 转换为 Application ImageTagRO
func ImageTagMapper(it *image_tags.ImageTag) ImageTagRO {
	if it == nil {
		return ImageTagRO{}
	}

	return ImageTagRO{
		ID:         it.ID(),
		ImageID:    it.ImageID(),
		Tag:        it.Tag(),
		Confidence: it.Confidence(),
		CreatedAt:  it.CreatedAt(),
	}
}
