package system_state

import (
	"time"

	"github.com/google/uuid"
)

type SystemState struct {
	state SystemStateState
}

type SystemStateState struct {
	ID                   uuid.UUID
	Initialized          bool
	InitializedAt        *time.Time
	InitializationVersion *string
	LastResetAt          *time.Time
	LastResetBy          *uuid.UUID
	ResetCount           int
	Metadata             map[string]interface{}
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (ss *SystemState) ID() uuid.UUID                    { return ss.state.ID }
func (ss *SystemState) Initialized() bool                { return ss.state.Initialized }
func (ss *SystemState) InitializedAt() *time.Time        { return ss.state.InitializedAt }
func (ss *SystemState) InitializationVersion() *string   { return ss.state.InitializationVersion }
func (ss *SystemState) LastResetAt() *time.Time          { return ss.state.LastResetAt }
func (ss *SystemState) LastResetBy() *uuid.UUID          { return ss.state.LastResetBy }
func (ss *SystemState) ResetCount() int                  { return ss.state.ResetCount }
func (ss *SystemState) Metadata() map[string]interface{} { return ss.state.Metadata }
func (ss *SystemState) CreatedAt() time.Time             { return ss.state.CreatedAt }
func (ss *SystemState) UpdatedAt() time.Time             { return ss.state.UpdatedAt }
