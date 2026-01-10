package results

import (
	"time"

	"nfxid/modules/system/domain/system_state"

	"github.com/google/uuid"
)

type SystemStateRO struct {
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

// SystemStateMapper 将 Domain SystemState 转换为 Application SystemStateRO
func SystemStateMapper(ss *system_state.SystemState) SystemStateRO {
	if ss == nil {
		return SystemStateRO{}
	}

	return SystemStateRO{
		ID:                   ss.ID(),
		Initialized:          ss.Initialized(),
		InitializedAt:        ss.InitializedAt(),
		InitializationVersion: ss.InitializationVersion(),
		LastResetAt:          ss.LastResetAt(),
		LastResetBy:          ss.LastResetBy(),
		ResetCount:           ss.ResetCount(),
		Metadata:             ss.Metadata(),
		CreatedAt:            ss.CreatedAt(),
		UpdatedAt:            ss.UpdatedAt(),
	}
}
