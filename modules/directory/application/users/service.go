package users

import (
	userDomain "nfxid/modules/directory/domain/users"
)

type Service struct {
	userRepo *userDomain.Repo
}

func NewService(
	userRepo *userDomain.Repo,
) *Service {
	return &Service{
		userRepo: userRepo,
	}
}
