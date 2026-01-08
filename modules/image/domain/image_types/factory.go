package image_types

import (
	"time"

	"github.com/google/uuid"
)

type NewImageTypeParams struct {
	Key         string
	Description *string
	MaxWidth    *int
	MaxHeight   *int
	AspectRatio *string
	IsSystem    bool
}

func NewImageType(p NewImageTypeParams) (*ImageType, error) {
	if err := validateImageTypeParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewImageTypeFromState(ImageTypeState{
		ID:          id,
		Key:         p.Key,
		Description: p.Description,
		MaxWidth:    p.MaxWidth,
		MaxHeight:   p.MaxHeight,
		AspectRatio: p.AspectRatio,
		IsSystem:    p.IsSystem,
		CreatedAt:   now,
		UpdatedAt:   now,
	}), nil
}

func NewImageTypeFromState(st ImageTypeState) *ImageType {
	return &ImageType{state: st}
}

func validateImageTypeParams(p NewImageTypeParams) error {
	if p.Key == "" {
		return ErrKeyRequired
	}
	return nil
}
