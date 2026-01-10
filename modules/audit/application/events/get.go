package events

import (
	"context"
	"time"
	eventResult "nfxid/modules/audit/application/events/results"
	eventDomain "nfxid/modules/audit/domain/events"

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

// GetEventsByActor 根据操作者获取事件列表
func (s *Service) GetEventsByActor(ctx context.Context, actorType eventDomain.ActorType, actorID uuid.UUID, startTime, endTime *time.Time) ([]eventResult.EventRO, error) {
	domainEntities, err := s.eventRepo.Get.ByActor(ctx, actorType, actorID, startTime, endTime)
	if err != nil {
		return nil, err
	}

	results := make([]eventResult.EventRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = eventResult.EventMapper(entity)
	}
	return results, nil
}
