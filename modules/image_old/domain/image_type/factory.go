package image_type

import (
	"time"

	"github.com/google/uuid"
)

func NewImageType(editable ImageTypeEditable) *ImageType {
	now := time.Now()
	return &ImageType{
		state: ImageTypeState{
			ID:          uuid.New(),
			Key:         editable.Key,
			Description: editable.Description,
			MaxWidth:    editable.MaxWidth,
			MaxHeight:   editable.MaxHeight,
			AspectRatio: editable.AspectRatio,
			IsSystem:    editable.IsSystem,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}
}

func NewImageTypeWithID(id uuid.UUID, editable ImageTypeEditable) *ImageType {
	now := time.Now()
	return &ImageType{
		state: ImageTypeState{
			ID:          id,
			Key:         editable.Key,
			Description: editable.Description,
			MaxWidth:    editable.MaxWidth,
			MaxHeight:   editable.MaxHeight,
			AspectRatio: editable.AspectRatio,
			IsSystem:    editable.IsSystem,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}
}

func NewImageTypeFromState(state ImageTypeState) *ImageType {
	return &ImageType{
		state: state,
	}
}

