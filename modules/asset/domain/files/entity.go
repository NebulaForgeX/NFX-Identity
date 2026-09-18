package files

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	state FileState
}

type FileState struct {
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

func (e *File) State() FileState { return e.state }

func (e *File) ID() uuid.UUID         { return e.state.ID }
func (e *File) FilePath() string      { return e.state.FilePath }
func (e *File) FileName() string      { return e.state.FileName }
func (e *File) FileSize() int64       { return e.state.FileSize }
func (e *File) MimeType() string      { return e.state.MimeType }
func (e *File) UploaderID() uuid.UUID { return e.state.UploaderID }
func (e *File) CreatedAt() time.Time  { return e.state.CreatedAt }
func (e *File) UpdatedAt() time.Time  { return e.state.UpdatedAt }
func (e *File) DeletedAt() *time.Time { return e.state.DeletedAt }

func NewFileFromState(st FileState) *File {
	return &File{state: st}
}
