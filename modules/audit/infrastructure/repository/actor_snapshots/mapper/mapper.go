package mapper

import (
	"encoding/json"
	"nfxid/enums"
	"nfxid/modules/audit/domain/actor_snapshots"
	"nfxid/modules/audit/infrastructure/rdb/models"

	"gorm.io/datatypes"
)

// ActorSnapshotDomainToModel 将 Domain ActorSnapshot 转换为 Model ActorSnapshot
func ActorSnapshotDomainToModel(as *actor_snapshots.ActorSnapshot) *models.ActorSnapshot {
	if as == nil {
		return nil
	}

	var snapshotData *datatypes.JSON
	if as.SnapshotData() != nil && len(as.SnapshotData()) > 0 {
		dataBytes, _ := json.Marshal(as.SnapshotData())
		jsonData := datatypes.JSON(dataBytes)
		snapshotData = &jsonData
	}

	return &models.ActorSnapshot{
		ID:           as.ID(),
		ActorType:    enums.AuditActorType(as.ActorType()),
		ActorID:      as.ActorID(),
		DisplayName:  as.DisplayName(),
		Email:        as.Email(),
		ClientName:   as.ClientName(),
		TenantID:     as.TenantID(),
		SnapshotAt:   as.SnapshotAt(),
		SnapshotData: snapshotData,
		CreatedAt:    as.CreatedAt(),
	}
}

// ActorSnapshotModelToDomain 将 Model ActorSnapshot 转换为 Domain ActorSnapshot
func ActorSnapshotModelToDomain(m *models.ActorSnapshot) *actor_snapshots.ActorSnapshot {
	if m == nil {
		return nil
	}

	var snapshotData map[string]interface{}
	if m.SnapshotData != nil {
		json.Unmarshal(*m.SnapshotData, &snapshotData)
	}

	state := actor_snapshots.ActorSnapshotState{
		ID:           m.ID,
		ActorType:    actor_snapshots.ActorType(m.ActorType),
		ActorID:      m.ActorID,
		DisplayName:  m.DisplayName,
		Email:        m.Email,
		ClientName:   m.ClientName,
		TenantID:     m.TenantID,
		SnapshotAt:   m.SnapshotAt,
		SnapshotData: snapshotData,
		CreatedAt:    m.CreatedAt,
	}

	return actor_snapshots.NewActorSnapshotFromState(state)
}

// ActorSnapshotModelToUpdates 将 Model ActorSnapshot 转换为更新字段映射
func ActorSnapshotModelToUpdates(m *models.ActorSnapshot) map[string]any {
	var snapshotData any
	if m.SnapshotData != nil {
		snapshotData = m.SnapshotData
	}

	return map[string]any{
		models.ActorSnapshotCols.ActorType:    m.ActorType,
		models.ActorSnapshotCols.ActorID:      m.ActorID,
		models.ActorSnapshotCols.DisplayName:  m.DisplayName,
		models.ActorSnapshotCols.Email:        m.Email,
		models.ActorSnapshotCols.ClientName:   m.ClientName,
		models.ActorSnapshotCols.TenantID:     m.TenantID,
		models.ActorSnapshotCols.SnapshotAt:   m.SnapshotAt,
		models.ActorSnapshotCols.SnapshotData: snapshotData,
	}
}
