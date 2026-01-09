package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/events"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/events/mapper"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByActor 根据 Actor 获取 Events，实现 events.Get 接口
func (h *Handler) ByActor(ctx context.Context, actorType events.ActorType, actorID uuid.UUID, startTime, endTime *time.Time) ([]*events.Event, error) {
	query := h.db.WithContext(ctx).
		Where("actor_type = ? AND actor_id = ?", actorType, actorID)
	
	if startTime != nil {
		query = query.Where("occurred_at >= ?", *startTime)
	}
	if endTime != nil {
		query = query.Where("occurred_at <= ?", *endTime)
	}
	
	var ms []models.Event
	if err := query.Find(&ms).Error; err != nil {
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
