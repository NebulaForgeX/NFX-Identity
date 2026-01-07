package profile_badge

import (
	profileBadgeDomain "nfxid/modules/auth/domain/profile_badge"
)

type Service struct {
	profileBadgeRepo  *profileBadgeDomain.Repo
	profileBadgeQuery *profileBadgeDomain.Query
}

func NewService(
	profileBadgeRepo *profileBadgeDomain.Repo,
	profileBadgeQuery *profileBadgeDomain.Query,
) *Service {
	return &Service{
		profileBadgeRepo:  profileBadgeRepo,
		profileBadgeQuery: profileBadgeQuery,
	}
}
