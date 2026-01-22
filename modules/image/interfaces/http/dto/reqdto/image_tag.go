package reqdto

import (
	imageTagAppCommands "nfxid/modules/image/application/image_tags/commands"

	"github.com/google/uuid"
)

type ImageTagCreateRequestDTO struct {
	ImageID    uuid.UUID  `json:"image_id" validate:"required,uuid"`
	Tag        string     `json:"tag" validate:"required"`
	Confidence *float64   `json:"confidence,omitempty"`
}

type ImageTagByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type ImageTagUpdateRequestDTO struct {
	Tag        string    `json:"tag" validate:"required"`
	Confidence *float64  `json:"confidence,omitempty"`
}

func (r *ImageTagCreateRequestDTO) ToCreateCmd() imageTagAppCommands.CreateImageTagCmd {
	return imageTagAppCommands.CreateImageTagCmd{
		ImageID:    r.ImageID,
		Tag:        r.Tag,
		Confidence: r.Confidence,
	}
}

func (r *ImageTagUpdateRequestDTO) ToUpdateCmd(imageTagID uuid.UUID) imageTagAppCommands.UpdateImageTagCmd {
	return imageTagAppCommands.UpdateImageTagCmd{
		ImageTagID: imageTagID,
		Tag:        r.Tag,
		Confidence: r.Confidence,
	}
}
