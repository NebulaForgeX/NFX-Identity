package reqdto

import (
	imageTypeAppCommands "nfxid/modules/image/application/image_types/commands"

	"github.com/google/uuid"
)

type ImageTypeCreateRequestDTO struct {
	Key         string  `json:"key" validate:"required"`
	Description *string `json:"description,omitempty"`
	MaxWidth    *int    `json:"max_width,omitempty"`
	MaxHeight   *int    `json:"max_height,omitempty"`
	AspectRatio *string `json:"aspect_ratio,omitempty"`
	IsSystem    bool    `json:"is_system"`
}

type ImageTypeByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type ImageTypeByKeyRequestDTO struct {
	Key string `params:"key" validate:"required"`
}

type ImageTypeUpdateRequestDTO struct {
	Key         string  `json:"key" validate:"required"`
	Description *string `json:"description,omitempty"`
	MaxWidth    *int    `json:"max_width,omitempty"`
	MaxHeight   *int    `json:"max_height,omitempty"`
	AspectRatio *string `json:"aspect_ratio,omitempty"`
}

func (r *ImageTypeCreateRequestDTO) ToCreateCmd() imageTypeAppCommands.CreateImageTypeCmd {
	return imageTypeAppCommands.CreateImageTypeCmd{
		Key:         r.Key,
		Description: r.Description,
		MaxWidth:    r.MaxWidth,
		MaxHeight:   r.MaxHeight,
		AspectRatio: r.AspectRatio,
		IsSystem:    r.IsSystem,
	}
}

func (r *ImageTypeUpdateRequestDTO) ToUpdateCmd(imageTypeID uuid.UUID) imageTypeAppCommands.UpdateImageTypeCmd {
	return imageTypeAppCommands.UpdateImageTypeCmd{
		ImageTypeID: imageTypeID,
		Key:         r.Key,
		Description: r.Description,
		MaxWidth:    r.MaxWidth,
		MaxHeight:   r.MaxHeight,
		AspectRatio: r.AspectRatio,
	}
}
