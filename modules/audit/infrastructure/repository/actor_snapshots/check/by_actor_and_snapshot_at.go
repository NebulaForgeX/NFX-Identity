package check

import (
	"context"
	"nfxid/modules/audit/domain/actor_snapshots"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"time"

	"github.com/google/uuid"
)

// ByActorAndSnapshotAt 根据 Actor 和 SnapshotAt 检查 ActorSnapshot 是否存在，实现 actor_snapshots.Check 接口
func (h *Handler) ByActorAndSnapshotAt(ctx context.Context, actorType actor_snapshots.ActorType, actorID uuid.UUID, snapshotAt time.Time) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.ActorSnapshot{}).
		Where("actor_type = ? AND actor_id = ? AND snapshot_at = ?", actorType, actorID, snapshotAt).
		Count(&count).Error
	return count > 0, err
}
