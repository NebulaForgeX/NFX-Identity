package user_badges

import (
	userBadgeDomain "nfxid/modules/directory/domain/user_badges"
)

type Service struct {
	userBadgeRepo *userBadgeDomain.Repo
}

func NewService(
	userBadgeRepo *userBadgeDomain.Repo,
) *Service {
	return &Service{
		userBadgeRepo: userBadgeRepo,
	}
}
