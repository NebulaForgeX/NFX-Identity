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

// LatestByActor 获取 Actor 的最新 ActorSnapshot，实现 actor_snapshots.Get 接口
func (h *Handler) LatestByActor(ctx context.Context, actorType actor_snapshots.ActorType, actorID uuid.UUID) (*actor_snapshots.ActorSnapshot, error) {
	var m models.ActorSnapshot
	if err := h.db.WithContext(ctx).
		Where("actor_type = ? AND actor_id = ?", actorType, actorID).
		Order("snapshot_at DESC").
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, actor_snapshots.ErrActorSnapshotNotFound
		}
		return nil, err
	}
	return mapper.ActorSnapshotModelToDomain(&m), nil
}
