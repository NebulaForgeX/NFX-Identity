package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/actor_snapshots"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/actor_snapshots/mapper"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByActorAndSnapshotAt 根据 Actor 和 SnapshotAt 获取 ActorSnapshot，实现 actor_snapshots.Get 接口
func (h *Handler) ByActorAndSnapshotAt(ctx context.Context, actorType actor_snapshots.ActorType, actorID uuid.UUID, snapshotAt time.Time) (*actor_snapshots.ActorSnapshot, error) {
	var m models.ActorSnapshot
	if err := h.db.WithContext(ctx).
		Where("actor_type = ? AND actor_id = ? AND snapshot_at = ?", actorType, actorID, snapshotAt).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, actor_snapshots.ErrActorSnapshotNotFound
		}
		return nil, err
	}
	return mapper.ActorSnapshotModelToDomain(&m), nil
}
