package profile_badge

import (
	profileBadgeQueries "nebulaid/modules/auth/application/profile_badge/queries"
	profileBadgeDomain "nebulaid/modules/auth/domain/profile_badge"
)

type Service struct {
	profileBadgeRepo  profileBadgeDomain.Repo
	profileBadgeQuery profileBadgeQueries.ProfileBadgeQuery
}

func NewService(
	profileBadgeRepo profileBadgeDomain.Repo,
	profileBadgeQuery profileBadgeQueries.ProfileBadgeQuery,
) *Service {
	return &Service{
		profileBadgeRepo:  profileBadgeRepo,
		profileBadgeQuery: profileBadgeQuery,
	}
}
