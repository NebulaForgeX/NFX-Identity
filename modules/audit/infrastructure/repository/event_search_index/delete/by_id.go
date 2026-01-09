package delete

import (
	"context"
	"nfxid/modules/audit/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 EventSearchIndex，实现 event_search_index.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.EventSearchIndex{}).Error
}
