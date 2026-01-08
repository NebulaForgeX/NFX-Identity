package image_tags

import (
	"time"

	"github.com/google/uuid"
)

type ImageTag struct {
	state ImageTagState
}

type ImageTagState struct {
	ID         uuid.UUID
	ImageID    uuid.UUID
	Tag        string
	Confidence *float64
	CreatedAt  time.Time
}

func (it *ImageTag) ID() uuid.UUID        { return it.state.ID }
func (it *ImageTag) ImageID() uuid.UUID    { return it.state.ImageID }
func (it *ImageTag) Tag() string           { return it.state.Tag }
func (it *ImageTag) Confidence() *float64 { return it.state.Confidence }
func (it *ImageTag) CreatedAt() time.Time { return it.state.CreatedAt }
