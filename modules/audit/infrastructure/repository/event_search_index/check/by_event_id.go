package check

import (
	"context"
	"nfxid/modules/audit/infrastructure/rdb/models"
)

// ByEventID 根据 EventID 检查 EventSearchIndex 是否存在，实现 event_search_index.Check 接口
func (h *Handler) ByEventID(ctx context.Context, eventID string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.EventSearchIndex{}).
		Where("event_id = ?", eventID).
		Count(&count).Error
	return count > 0, err
}
