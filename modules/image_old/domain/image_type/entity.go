package image_type

import (
	"time"

	"github.com/google/uuid"
)

type ImageType struct {
	state ImageTypeState
}

type ImageTypeState struct {
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

type ImageTypeEditable struct {
	Key         string
	Description *string
	MaxWidth    *int
	MaxHeight   *int
	AspectRatio *string
	IsSystem    bool
}

func (it *ImageType) ID() uuid.UUID        { return it.state.ID }
func (it *ImageType) Key() string          { return it.state.Key }
func (it *ImageType) Description() *string { return it.state.Description }
func (it *ImageType) MaxWidth() *int       { return it.state.MaxWidth }
func (it *ImageType) MaxHeight() *int      { return it.state.MaxHeight }
func (it *ImageType) AspectRatio() *string { return it.state.AspectRatio }
func (it *ImageType) IsSystem() bool       { return it.state.IsSystem }
func (it *ImageType) CreatedAt() time.Time { return it.state.CreatedAt }
func (it *ImageType) UpdatedAt() time.Time { return it.state.UpdatedAt }
