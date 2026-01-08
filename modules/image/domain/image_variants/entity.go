package image_variants

import (
	"time"

	"github.com/google/uuid"
)

type ImageVariant struct {
	state ImageVariantState
}

type ImageVariantState struct {
	ID          uuid.UUID
	ImageID     uuid.UUID
	VariantKey  string
	Width       *int
	Height      *int
	Size        *int64
	MimeType    *string
	StoragePath string
	URL         *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (iv *ImageVariant) ID() uuid.UUID          { return iv.state.ID }
func (iv *ImageVariant) ImageID() uuid.UUID      { return iv.state.ImageID }
func (iv *ImageVariant) VariantKey() string      { return iv.state.VariantKey }
func (iv *ImageVariant) Width() *int            { return iv.state.Width }
func (iv *ImageVariant) Height() *int           { return iv.state.Height }
func (iv *ImageVariant) Size() *int64           { return iv.state.Size }
func (iv *ImageVariant) MimeType() *string      { return iv.state.MimeType }
func (iv *ImageVariant) StoragePath() string    { return iv.state.StoragePath }
func (iv *ImageVariant) URL() *string           { return iv.state.URL }
func (iv *ImageVariant) CreatedAt() time.Time   { return iv.state.CreatedAt }
func (iv *ImageVariant) UpdatedAt() time.Time   { return iv.state.UpdatedAt }
