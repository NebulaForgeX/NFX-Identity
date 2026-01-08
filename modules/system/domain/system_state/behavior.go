package system_state

import (
	"time"

	"github.com/google/uuid"
)

func (ss *SystemState) Initialize(version string) error {
	if ss.Initialized() {
		return ErrAlreadyInitialized
	}

	now := time.Now().UTC()
	ss.state.Initialized = true
	ss.state.InitializedAt = &now
	if version != "" {
		ss.state.InitializationVersion = &version
	}
	ss.state.UpdatedAt = now
	return nil
}

func (ss *SystemState) Reset(resetBy uuid.UUID) error {
	now := time.Now().UTC()
	ss.state.Initialized = false
	ss.state.InitializedAt = nil
	ss.state.InitializationVersion = nil
	ss.state.LastResetAt = &now
	ss.state.LastResetBy = &resetBy
	ss.state.ResetCount++
	ss.state.UpdatedAt = now
	return nil
}

func (ss *SystemState) UpdateMetadata(metadata map[string]interface{}) error {
	if metadata == nil {
		return nil
	}
	ss.state.Metadata = metadata
	ss.state.UpdatedAt = time.Now().UTC()
	return nil
}
