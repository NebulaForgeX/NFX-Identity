package audios

import (
	"time"

	"github.com/google/uuid"
)

type Audio struct {
	state AudioState
}

type AudioState struct {
	ID              uuid.UUID
	FilePath        string
	FileName        string
	FileSize        int64
	MimeType        string
	DurationSeconds *float64
	UploaderID      uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

func (e *Audio) State() AudioState { return e.state }

func (e *Audio) ID() uuid.UUID             { return e.state.ID }
func (e *Audio) FilePath() string          { return e.state.FilePath }
func (e *Audio) FileName() string          { return e.state.FileName }
func (e *Audio) FileSize() int64           { return e.state.FileSize }
func (e *Audio) MimeType() string          { return e.state.MimeType }
func (e *Audio) DurationSeconds() *float64 { return e.state.DurationSeconds }
func (e *Audio) UploaderID() uuid.UUID     { return e.state.UploaderID }
func (e *Audio) CreatedAt() time.Time      { return e.state.CreatedAt }
func (e *Audio) UpdatedAt() time.Time      { return e.state.UpdatedAt }
func (e *Audio) DeletedAt() *time.Time     { return e.state.DeletedAt }

func NewAudioFromState(st AudioState) *Audio {
	return &Audio{state: st}
}
