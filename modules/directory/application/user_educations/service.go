package user_educations

import (
	userEducationDomain "nfxidentity/modules/directory/domain/user_educations"
)

type Service struct {
	userEducationRepo *userEducationDomain.Repo
}

func NewService(
	userEducationRepo *userEducationDomain.Repo,
) *Service {
	return &Service{
		userEducationRepo: userEducationRepo,
	}
}
