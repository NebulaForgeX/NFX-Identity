package update

import (
	"context"
	"nfxid/modules/audit/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// Status 更新 EventRetentionPolicy 的 Status，实现 event_retention_policies.Update 接口
func (h *Handler) Status(ctx context.Context, id uuid.UUID, status string) error {
	return h.db.WithContext(ctx).
		Model(&models.EventRetentionPolicy{}).
		Where("id = ?", id).
		Update(models.EventRetentionPolicyCols.Status, status).Error
}
