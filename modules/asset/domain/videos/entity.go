package videos

import (
	"time"

	"github.com/google/uuid"
)

type Video struct {
	state VideoState
}

type VideoState struct {
	ID              uuid.UUID
	FilePath        string
	FileName        string
	FileSize        int64
	MimeType        string
	DurationSeconds *float64
	Width           *int
	Height          *int
	UploaderID      uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

func (e *Video) State() VideoState { return e.state }

func (e *Video) ID() uuid.UUID             { return e.state.ID }
func (e *Video) FilePath() string          { return e.state.FilePath }
func (e *Video) FileName() string          { return e.state.FileName }
func (e *Video) FileSize() int64           { return e.state.FileSize }
func (e *Video) MimeType() string          { return e.state.MimeType }
func (e *Video) DurationSeconds() *float64 { return e.state.DurationSeconds }
func (e *Video) Width() *int               { return e.state.Width }
func (e *Video) Height() *int              { return e.state.Height }
func (e *Video) UploaderID() uuid.UUID     { return e.state.UploaderID }
func (e *Video) CreatedAt() time.Time      { return e.state.CreatedAt }
func (e *Video) UpdatedAt() time.Time      { return e.state.UpdatedAt }
func (e *Video) DeletedAt() *time.Time     { return e.state.DeletedAt }

func NewVideoFromState(st VideoState) *Video {
	return &Video{state: st}
}
