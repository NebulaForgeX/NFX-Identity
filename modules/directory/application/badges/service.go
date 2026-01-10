package badges

import (
	badgeDomain "nfxid/modules/directory/domain/badges"
)

type Service struct {
	badgeRepo *badgeDomain.Repo
}

func NewService(
	badgeRepo *badgeDomain.Repo,
) *Service {
	return &Service{
		badgeRepo: badgeRepo,
	}
}
