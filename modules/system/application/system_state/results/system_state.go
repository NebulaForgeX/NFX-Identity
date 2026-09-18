package results

import (
	"time"

	"nfxidentity/modules/system/domain/system_state"

	"github.com/google/uuid"
)

type SystemStateRO struct {
	ID                    uuid.UUID              `json:"id"`
	Initialized           bool                   `json:"initialized"`
	InitializedAt         *time.Time             `json:"initialized_at,omitempty"`
	InitializationVersion *string                `json:"initialization_version,omitempty"`
	LastResetAt           *time.Time             `json:"last_reset_at,omitempty"`
	LastResetBy           *uuid.UUID             `json:"last_reset_by,omitempty"`
	ResetCount            int                    `json:"reset_count"`
	Metadata              map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt             time.Time              `json:"created_at"`
	UpdatedAt             time.Time              `json:"updated_at"`
}

// SystemStateMapper 将 Domain SystemState 转换为 Application SystemStateRO
func SystemStateMapper(ss *system_state.SystemState) SystemStateRO {
	if ss == nil {
		return SystemStateRO{}
	}

	return SystemStateRO{
		ID:                    ss.ID(),
		Initialized:           ss.Initialized(),
		InitializedAt:         ss.InitializedAt(),
		InitializationVersion: ss.InitializationVersion(),
		LastResetAt:           ss.LastResetAt(),
		LastResetBy:           ss.LastResetBy(),
		ResetCount:            ss.ResetCount(),
		Metadata:              ss.Metadata(),
		CreatedAt:             ss.CreatedAt(),
		UpdatedAt:             ss.UpdatedAt(),
	}
}
