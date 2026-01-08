package image_tags

import (
	"time"

	"github.com/google/uuid"
)

type NewImageTagParams struct {
	ImageID    uuid.UUID
	Tag        string
	Confidence *float64
}

func NewImageTag(p NewImageTagParams) (*ImageTag, error) {
	if err := validateImageTagParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewImageTagFromState(ImageTagState{
		ID:         id,
		ImageID:    p.ImageID,
		Tag:        p.Tag,
		Confidence: p.Confidence,
		CreatedAt:  now,
	}), nil
}

func NewImageTagFromState(st ImageTagState) *ImageTag {
	return &ImageTag{state: st}
}

func validateImageTagParams(p NewImageTagParams) error {
	if p.ImageID == uuid.Nil {
		return ErrImageIDRequired
	}
	if p.Tag == "" {
		return ErrTagRequired
	}
	return nil
}
