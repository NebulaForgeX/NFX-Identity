package user_occupations

import (
	userOccupationDomain "nfxid/modules/directory/domain/user_occupations"
)

type Service struct {
	userOccupationRepo *userOccupationDomain.Repo
}

func NewService(
	userOccupationRepo *userOccupationDomain.Repo,
) *Service {
	return &Service{
		userOccupationRepo: userOccupationRepo,
	}
}
