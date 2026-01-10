package user_preferences

import (
	userPreferenceDomain "nfxid/modules/directory/domain/user_preferences"
)

type Service struct {
	userPreferenceRepo *userPreferenceDomain.Repo
}

func NewService(
	userPreferenceRepo *userPreferenceDomain.Repo,
) *Service {
	return &Service{
		userPreferenceRepo: userPreferenceRepo,
	}
}
