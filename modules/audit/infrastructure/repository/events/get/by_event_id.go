package get

import (
	"context"
	"errors"
	auditErr "nfxidentity/errors/src/audit"
	"nfxidentity/modules/audit/domain/events"
	"nfxidentity/modules/audit/infrastructure/rdb/models"
	"nfxidentity/modules/audit/infrastructure/repository/events/mapper"

	"gorm.io/gorm"
)

// ByEventID 根据 EventID 获取 Event，实现 events.Get 接口
func (h *Handler) ByEventID(ctx context.Context, eventID string) (*events.Event, error) {
	var m models.Event
	if err := h.db.WithContext(ctx).Where("event_id = ?", eventID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, auditErr.ErrEventNotFound
		}
		return nil, err
	}
	return mapper.EventModelToDomain(&m), nil
}
