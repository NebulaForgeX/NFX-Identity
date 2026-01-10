package user_phones

import (
	userPhoneDomain "nfxid/modules/directory/domain/user_phones"
)

type Service struct {
	userPhoneRepo *userPhoneDomain.Repo
}

func NewService(
	userPhoneRepo *userPhoneDomain.Repo,
) *Service {
	return &Service{
		userPhoneRepo: userPhoneRepo,
	}
}
