package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/events"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/events/mapper"

	"gorm.io/gorm"
)

// ByRequestID 根据 RequestID 获取 Events，实现 events.Get 接口
func (h *Handler) ByRequestID(ctx context.Context, requestID string) ([]*events.Event, error) {
	var ms []models.Event
	if err := h.db.WithContext(ctx).Where("request_id = ?", requestID).Find(&ms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*events.Event{}, nil
		}
		return nil, err
	}
	
	result := make([]*events.Event, len(ms))
	for i := range ms {
		result[i] = mapper.EventModelToDomain(&ms[i])
	}
	return result, nil
}
