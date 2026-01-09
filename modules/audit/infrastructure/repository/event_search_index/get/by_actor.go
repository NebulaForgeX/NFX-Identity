package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/event_search_index"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/event_search_index/mapper"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByActor 根据 Actor 获取 EventSearchIndexes，实现 event_search_index.Get 接口
func (h *Handler) ByActor(ctx context.Context, actorType event_search_index.ActorType, actorID uuid.UUID, startTime, endTime *time.Time) ([]*event_search_index.EventSearchIndex, error) {
	query := h.db.WithContext(ctx).
		Where("actor_type = ? AND actor_id = ?", actorType, actorID)
	
	if startTime != nil {
		query = query.Where("occurred_at >= ?", *startTime)
	}
	if endTime != nil {
		query = query.Where("occurred_at <= ?", *endTime)
	}
	
	var ms []models.EventSearchIndex
	if err := query.Find(&ms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*event_search_index.EventSearchIndex{}, nil
		}
		return nil, err
	}
	
	result := make([]*event_search_index.EventSearchIndex, len(ms))
	for i := range ms {
		result[i] = mapper.EventSearchIndexModelToDomain(&ms[i])
	}
	return result, nil
}
