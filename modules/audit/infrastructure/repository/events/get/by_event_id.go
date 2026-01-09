package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/events"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/events/mapper"

	"gorm.io/gorm"
)

// ByEventID 根据 EventID 获取 Event，实现 events.Get 接口
func (h *Handler) ByEventID(ctx context.Context, eventID string) (*events.Event, error) {
	var m models.Event
	if err := h.db.WithContext(ctx).Where("event_id = ?", eventID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, events.ErrEventNotFound
		}
		return nil, err
	}
	return mapper.EventModelToDomain(&m), nil
}
