package mapper

import (
	"encoding/json"
	"nfxid/modules/system/domain/system_state"
	"nfxid/modules/system/infrastructure/rdb/models"

	"gorm.io/datatypes"
)

// SystemStateDomainToModel 将 Domain SystemState 转换为 Model SystemState
func SystemStateDomainToModel(ss *system_state.SystemState) *models.SystemState {
	if ss == nil {
		return nil
	}

	var metadata *datatypes.JSON
	if ss.Metadata() != nil && len(ss.Metadata()) > 0 {
		metaBytes, _ := json.Marshal(ss.Metadata())
		jsonData := datatypes.JSON(metaBytes)
		metadata = &jsonData
	}

	return &models.SystemState{
		ID:                    ss.ID(),
		Initialized:           ss.Initialized(),
		InitializedAt:         ss.InitializedAt(),
		InitializationVersion: ss.InitializationVersion(),
		LastResetAt:           ss.LastResetAt(),
		LastResetBy:           ss.LastResetBy(),
		ResetCount:            ss.ResetCount(),
		Metadata:              metadata,
		CreatedAt:             ss.CreatedAt(),
		UpdatedAt:             ss.UpdatedAt(),
	}
}

// SystemStateModelToDomain 将 Model SystemState 转换为 Domain SystemState
func SystemStateModelToDomain(m *models.SystemState) *system_state.SystemState {
	if m == nil {
		return nil
	}

	var metadata map[string]interface{}
	if m.Metadata != nil {
		json.Unmarshal(*m.Metadata, &metadata)
	}

	state := system_state.SystemStateState{
		ID:                    m.ID,
		Initialized:           m.Initialized,
		InitializedAt:         m.InitializedAt,
		InitializationVersion: m.InitializationVersion,
		LastResetAt:           m.LastResetAt,
		LastResetBy:           m.LastResetBy,
		ResetCount:            m.ResetCount,
		Metadata:              metadata,
		CreatedAt:             m.CreatedAt,
		UpdatedAt:             m.UpdatedAt,
	}

	return system_state.NewSystemStateFromState(state)
}

// SystemStateModelToUpdates 将 Model SystemState 转换为更新字段映射
func SystemStateModelToUpdates(m *models.SystemState) map[string]any {
	var metadata any
	if m.Metadata != nil {
		metadata = m.Metadata
	}

	return map[string]any{
		models.SystemStateCols.Initialized:           m.Initialized,
		models.SystemStateCols.InitializedAt:         m.InitializedAt,
		models.SystemStateCols.InitializationVersion: m.InitializationVersion,
		models.SystemStateCols.LastResetAt:           m.LastResetAt,
		models.SystemStateCols.LastResetBy:           m.LastResetBy,
		models.SystemStateCols.ResetCount:            m.ResetCount,
		models.SystemStateCols.Metadata:              metadata,
		models.SystemStateCols.UpdatedAt:            m.UpdatedAt,
	}
}
