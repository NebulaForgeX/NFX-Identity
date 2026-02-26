package user_profiles

import (
	userProfileDomain "nfxidentity/modules/directory/domain/user_profiles"
)

type Service struct {
	userProfileRepo *userProfileDomain.Repo
}

func NewService(
	userProfileRepo *userProfileDomain.Repo,
) *Service {
	return &Service{
		userProfileRepo: userProfileRepo,
	}
}
