package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/actor_snapshots"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/actor_snapshots/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByActor 根据 Actor 获取 ActorSnapshots，实现 actor_snapshots.Get 接口
func (h *Handler) ByActor(ctx context.Context, actorType actor_snapshots.ActorType, actorID uuid.UUID) ([]*actor_snapshots.ActorSnapshot, error) {
	var ms []models.ActorSnapshot
	if err := h.db.WithContext(ctx).
		Where("actor_type = ? AND actor_id = ?", actorType, actorID).
		Order("snapshot_at DESC").
		Find(&ms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*actor_snapshots.ActorSnapshot{}, nil
		}
		return nil, err
	}
	
	result := make([]*actor_snapshots.ActorSnapshot, len(ms))
	for i := range ms {
		result[i] = mapper.ActorSnapshotModelToDomain(&ms[i])
	}
	return result, nil
}
