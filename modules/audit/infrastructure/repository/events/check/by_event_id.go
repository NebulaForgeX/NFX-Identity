package check

import (
	"context"
	"nfxid/modules/audit/infrastructure/rdb/models"
)

// ByEventID 根据 EventID 检查 Event 是否存在，实现 events.Check 接口
func (h *Handler) ByEventID(ctx context.Context, eventID string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.Event{}).
		Where("event_id = ?", eventID).
		Count(&count).Error
	return count > 0, err
}
