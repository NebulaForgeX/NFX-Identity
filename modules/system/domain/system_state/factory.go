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
	
	// 如果未初始化，InitializedAt 应该为 nil
	// 如果已初始化但没有提供 InitializedAt，使用当前时间
	var initializedAt *time.Time
	if p.Initialized {
		if p.InitializedAt != nil {
			initializedAt = p.InitializedAt
		} else {
			initializedAt = &now
		}
	} else {
		// 未初始化时，InitializedAt 必须为 nil
		initializedAt = nil
	}

	return NewSystemStateFromState(SystemStateState{
		ID:                   id,
		Initialized:          p.Initialized,
		InitializedAt:        initializedAt,
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
