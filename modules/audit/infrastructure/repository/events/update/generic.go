package update

import (
	"context"
	"nfxid/modules/audit/domain/events"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/events/mapper"
)

// Generic 通用更新 Event，实现 events.Update 接口
func (h *Handler) Generic(ctx context.Context, e *events.Event) error {
	m := mapper.EventDomainToModel(e)
	updates := mapper.EventModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Event{}).
		Where("id = ?", e.ID()).
		Updates(updates).Error
}
