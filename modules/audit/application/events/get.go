package events

import (
	"context"
	eventResult "nfxid/modules/audit/application/events/results"

	"github.com/google/uuid"
)

// GetEvent 根据ID获取事件
func (s *Service) GetEvent(ctx context.Context, eventID uuid.UUID) (eventResult.EventRO, error) {
	domainEntity, err := s.eventRepo.Get.ByID(ctx, eventID)
	if err != nil {
		return eventResult.EventRO{}, err
	}
	return eventResult.EventMapper(domainEntity), nil
}

// GetEventByEventID 根据EventID获取事件
func (s *Service) GetEventByEventID(ctx context.Context, eventID string) (eventResult.EventRO, error) {
	domainEntity, err := s.eventRepo.Get.ByEventID(ctx, eventID)
	if err != nil {
		return eventResult.EventRO{}, err
	}
	return eventResult.EventMapper(domainEntity), nil
}
