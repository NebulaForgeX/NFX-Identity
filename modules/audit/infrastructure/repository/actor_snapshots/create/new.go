package create

import (
	"context"
	"nfxid/modules/audit/domain/actor_snapshots"
	"nfxid/modules/audit/infrastructure/repository/actor_snapshots/mapper"
)

// New 创建新的 ActorSnapshot，实现 actor_snapshots.Create 接口
func (h *Handler) New(ctx context.Context, as *actor_snapshots.ActorSnapshot) error {
	m := mapper.ActorSnapshotDomainToModel(as)
	return h.db.WithContext(ctx).Create(&m).Error
}
