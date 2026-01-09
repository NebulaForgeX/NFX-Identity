package mapper

import (
	"nfxid/modules/audit/domain/hash_chain_checkpoints"
	"nfxid/modules/audit/infrastructure/rdb/models"
)

// HashChainCheckpointDomainToModel 将 Domain HashChainCheckpoint 转换为 Model HashChainCheckpoint
func HashChainCheckpointDomainToModel(hcc *hash_chain_checkpoints.HashChainCheckpoint) *models.HashChainCheckpoint {
	if hcc == nil {
		return nil
	}

	return &models.HashChainCheckpoint{
		ID:                 hcc.ID(),
		CheckpointID:       hcc.CheckpointID(),
		TenantID:           hcc.TenantID(),
		PartitionDate:      hcc.PartitionDate(),
		CheckpointHash:     hcc.CheckpointHash(),
		PrevCheckpointHash: hcc.PrevCheckpointHash(),
		EventCount:         hcc.EventCount(),
		FirstEventID:       hcc.FirstEventID(),
		LastEventID:        hcc.LastEventID(),
		CreatedAt:          hcc.CreatedAt(),
		CreatedBy:          hcc.CreatedBy(),
	}
}

// HashChainCheckpointModelToDomain 将 Model HashChainCheckpoint 转换为 Domain HashChainCheckpoint
func HashChainCheckpointModelToDomain(m *models.HashChainCheckpoint) *hash_chain_checkpoints.HashChainCheckpoint {
	if m == nil {
		return nil
	}

	state := hash_chain_checkpoints.HashChainCheckpointState{
		ID:                 m.ID,
		CheckpointID:       m.CheckpointID,
		TenantID:           m.TenantID,
		PartitionDate:      m.PartitionDate,
		CheckpointHash:     m.CheckpointHash,
		PrevCheckpointHash: m.PrevCheckpointHash,
		EventCount:         m.EventCount,
		FirstEventID:       m.FirstEventID,
		LastEventID:        m.LastEventID,
		CreatedAt:          m.CreatedAt,
		CreatedBy:          m.CreatedBy,
	}

	return hash_chain_checkpoints.NewHashChainCheckpointFromState(state)
}

// HashChainCheckpointModelToUpdates 将 Model HashChainCheckpoint 转换为更新字段映射
func HashChainCheckpointModelToUpdates(m *models.HashChainCheckpoint) map[string]any {
	return map[string]any{
		models.HashChainCheckpointCols.CheckpointID:       m.CheckpointID,
		models.HashChainCheckpointCols.TenantID:           m.TenantID,
		models.HashChainCheckpointCols.PartitionDate:      m.PartitionDate,
		models.HashChainCheckpointCols.CheckpointHash:     m.CheckpointHash,
		models.HashChainCheckpointCols.PrevCheckpointHash: m.PrevCheckpointHash,
		models.HashChainCheckpointCols.EventCount:         m.EventCount,
		models.HashChainCheckpointCols.FirstEventID:       m.FirstEventID,
		models.HashChainCheckpointCols.LastEventID:        m.LastEventID,
		models.HashChainCheckpointCols.CreatedBy:          m.CreatedBy,
	}
}
