package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/event_search_index"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/event_search_index/mapper"

	"gorm.io/gorm"
)

// ByEventID 根据 EventID 获取 EventSearchIndex，实现 event_search_index.Get 接口
func (h *Handler) ByEventID(ctx context.Context, eventID string) (*event_search_index.EventSearchIndex, error) {
	var m models.EventSearchIndex
	if err := h.db.WithContext(ctx).Where("event_id = ?", eventID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, event_search_index.ErrEventSearchIndexNotFound
		}
		return nil, err
	}
	return mapper.EventSearchIndexModelToDomain(&m), nil
}
