package images

import (
	"time"

	"github.com/google/uuid"
)

type Image struct{ state ImageState }

type ImageState struct {
	ID         uuid.UUID
	FilePath   string
	FileName   string
	FileSize   int64
	MimeType   string
	UploaderID uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func NewFromState(st ImageState) *Image { return &Image{state: st} }
func (i *Image) ID() uuid.UUID          { return i.state.ID }
func (i *Image) FilePath() string       { return i.state.FilePath }
func (i *Image) FileName() string       { return i.state.FileName }
func (i *Image) FileSize() int64        { return i.state.FileSize }
func (i *Image) MimeType() string       { return i.state.MimeType }
func (i *Image) UploaderID() uuid.UUID  { return i.state.UploaderID }
func (i *Image) CreatedAt() time.Time   { return i.state.CreatedAt }
func (i *Image) UpdatedAt() time.Time   { return i.state.UpdatedAt }
func (i *Image) DeletedAt() *time.Time  { return i.state.DeletedAt }
func (i *Image) State() ImageState      { return i.state }
