package event_search_index

import (
	"context"
	eventSearchIndexResult "nfxid/modules/audit/application/event_search_index/results"

	"github.com/google/uuid"
)

// GetEventSearchIndex 根据ID获取事件搜索索引
func (s *Service) GetEventSearchIndex(ctx context.Context, eventSearchIndexID uuid.UUID) (eventSearchIndexResult.EventSearchIndexRO, error) {
	domainEntity, err := s.eventSearchIndexRepo.Get.ByID(ctx, eventSearchIndexID)
	if err != nil {
		return eventSearchIndexResult.EventSearchIndexRO{}, err
	}
	return eventSearchIndexResult.EventSearchIndexMapper(domainEntity), nil
}
