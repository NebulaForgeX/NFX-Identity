package update

import (
	"context"
	"nfxid/modules/audit/domain/event_search_index"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/event_search_index/mapper"
)

// Generic 通用更新 EventSearchIndex，实现 event_search_index.Update 接口
func (h *Handler) Generic(ctx context.Context, esi *event_search_index.EventSearchIndex) error {
	m := mapper.EventSearchIndexDomainToModel(esi)
	updates := mapper.EventSearchIndexModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.EventSearchIndex{}).
		Where("id = ?", esi.ID()).
		Updates(updates).Error
}
