package events

import (
	eventDomain "nfxid/modules/audit/domain/events"
)

type Service struct {
	eventRepo *eventDomain.Repo
}

func NewService(
	eventRepo *eventDomain.Repo,
) *Service {
	return &Service{
		eventRepo: eventRepo,
	}
}
