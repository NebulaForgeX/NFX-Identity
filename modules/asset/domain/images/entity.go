package images

import (
	"time"

	"github.com/google/uuid"
)

// Image 对应 asset."Images" 聚合根（与 infrastructure/rdb/models.Image 列一致）。
type Image struct {
	state ImageState
}

type ImageState struct {
	ID         uuid.UUID
	FilePath   string
	FileName   string
	FileSize   int64
	MimeType   string
	Width      *int
	Height     *int
	AltText    *string
	UploaderID uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func (e *Image) State() ImageState { return e.state }

func (e *Image) ID() uuid.UUID         { return e.state.ID }
func (e *Image) FilePath() string      { return e.state.FilePath }
func (e *Image) FileName() string      { return e.state.FileName }
func (e *Image) FileSize() int64       { return e.state.FileSize }
func (e *Image) MimeType() string      { return e.state.MimeType }
func (e *Image) Width() *int           { return e.state.Width }
func (e *Image) Height() *int          { return e.state.Height }
func (e *Image) AltText() *string      { return e.state.AltText }
func (e *Image) UploaderID() uuid.UUID { return e.state.UploaderID }
func (e *Image) CreatedAt() time.Time  { return e.state.CreatedAt }
func (e *Image) UpdatedAt() time.Time  { return e.state.UpdatedAt }
func (e *Image) DeletedAt() *time.Time { return e.state.DeletedAt }

func NewImageFromState(st ImageState) *Image {
	return &Image{state: st}
}
