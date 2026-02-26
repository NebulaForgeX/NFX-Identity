package update

import (
	"context"
	"nfxidentity/modules/audit/domain/actor_snapshots"
	"nfxidentity/modules/audit/infrastructure/rdb/models"
	"nfxidentity/modules/audit/infrastructure/repository/actor_snapshots/mapper"
)

// Generic 通用更新 ActorSnapshot，实现 actor_snapshots.Update 接口
func (h *Handler) Generic(ctx context.Context, as *actor_snapshots.ActorSnapshot) error {
	m := mapper.ActorSnapshotDomainToModel(as)
	updates := mapper.ActorSnapshotModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.ActorSnapshot{}).
		Where("id = ?", as.ID()).
		Updates(updates).Error
}
