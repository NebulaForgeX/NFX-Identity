package delete

import (
	"context"
	"nfxid/modules/audit/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 EventRetentionPolicy，实现 event_retention_policies.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.EventRetentionPolicy{}).Error
}
