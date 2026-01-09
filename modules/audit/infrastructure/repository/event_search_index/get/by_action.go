package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/event_search_index"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/event_search_index/mapper"
	"time"

	"gorm.io/gorm"
)

// ByAction 根据 Action 获取 EventSearchIndexes，实现 event_search_index.Get 接口
func (h *Handler) ByAction(ctx context.Context, action string, startTime, endTime *time.Time) ([]*event_search_index.EventSearchIndex, error) {
	query := h.db.WithContext(ctx).Where("action = ?", action)
	
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
