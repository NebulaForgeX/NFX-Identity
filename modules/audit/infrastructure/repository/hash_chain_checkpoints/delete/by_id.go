package delete

import (
	"context"
	"nfxid/modules/audit/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 HashChainCheckpoint，实现 hash_chain_checkpoints.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.HashChainCheckpoint{}).Error
}
