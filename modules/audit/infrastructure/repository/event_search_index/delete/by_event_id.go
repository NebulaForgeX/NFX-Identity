package delete

import (
	"context"
	"nfxid/modules/audit/infrastructure/rdb/models"
)

// ByEventID 根据 EventID 删除 EventSearchIndex，实现 event_search_index.Delete 接口
func (h *Handler) ByEventID(ctx context.Context, eventID string) error {
	return h.db.WithContext(ctx).
		Where("event_id = ?", eventID).
		Delete(&models.EventSearchIndex{}).Error
}
