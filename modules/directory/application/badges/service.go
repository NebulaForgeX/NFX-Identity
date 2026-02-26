package badges

import (
	badgeDomain "nfxidentity/modules/directory/domain/badges"
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
