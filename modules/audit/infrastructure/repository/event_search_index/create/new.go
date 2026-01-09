package create

import (
	"context"
	"nfxid/modules/audit/domain/event_search_index"
	"nfxid/modules/audit/infrastructure/repository/event_search_index/mapper"
)

// New 创建新的 EventSearchIndex，实现 event_search_index.Create 接口
func (h *Handler) New(ctx context.Context, esi *event_search_index.EventSearchIndex) error {
	m := mapper.EventSearchIndexDomainToModel(esi)
	return h.db.WithContext(ctx).Create(&m).Error
}
