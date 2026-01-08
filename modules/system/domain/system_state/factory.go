package system_state

import (
	"time"

	"github.com/google/uuid"
)

type NewSystemStateParams struct {
	Initialized           bool
	InitializedAt         *time.Time
	InitializationVersion *string
	LastResetAt           *time.Time
	LastResetBy           *uuid.UUID
	ResetCount            int
	Metadata              map[string]interface{}
}

func NewSystemState(p NewSystemStateParams) (*SystemState, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	initializedAt := now
	if p.InitializedAt != nil {
		initializedAt = *p.InitializedAt
	}

	return NewSystemStateFromState(SystemStateState{
		ID:                   id,
		Initialized:          p.Initialized,
		InitializedAt:        &initializedAt,
		InitializationVersion: p.InitializationVersion,
		LastResetAt:          p.LastResetAt,
		LastResetBy:          p.LastResetBy,
		ResetCount:           p.ResetCount,
		Metadata:             p.Metadata,
		CreatedAt:            now,
		UpdatedAt:            now,
	}), nil
}

func NewSystemStateFromState(st SystemStateState) *SystemState {
	return &SystemState{state: st}
}
