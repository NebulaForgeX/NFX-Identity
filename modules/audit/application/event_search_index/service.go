package event_search_index

import (
	eventSearchIndexDomain "nfxid/modules/audit/domain/event_search_index"
)

type Service struct {
	eventSearchIndexRepo *eventSearchIndexDomain.Repo
}

func NewService(
	eventSearchIndexRepo *eventSearchIndexDomain.Repo,
) *Service {
	return &Service{
		eventSearchIndexRepo: eventSearchIndexRepo,
	}
}
