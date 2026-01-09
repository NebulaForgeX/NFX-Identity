package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/events"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/events/mapper"

	"gorm.io/gorm"
)

// ByTraceID 根据 TraceID 获取 Events，实现 events.Get 接口
func (h *Handler) ByTraceID(ctx context.Context, traceID string) ([]*events.Event, error) {
	var ms []models.Event
	if err := h.db.WithContext(ctx).Where("trace_id = ?", traceID).Find(&ms).Error; err != nil {
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
