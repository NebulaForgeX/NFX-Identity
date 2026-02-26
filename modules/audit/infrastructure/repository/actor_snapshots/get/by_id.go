package get

import (
	"context"
	"errors"
	auditErr "nfxidentity/errors/src/audit"
	"nfxidentity/modules/audit/domain/actor_snapshots"
	"nfxidentity/modules/audit/infrastructure/rdb/models"
	"nfxidentity/modules/audit/infrastructure/repository/actor_snapshots/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 ActorSnapshot，实现 actor_snapshots.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*actor_snapshots.ActorSnapshot, error) {
	var m models.ActorSnapshot
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, auditErr.ErrActorSnapshotNotFound
		}
		return nil, err
	}
	return mapper.ActorSnapshotModelToDomain(&m), nil
}
