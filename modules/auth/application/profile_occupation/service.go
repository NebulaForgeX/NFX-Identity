package profile_occupation

import (
	occupationDomain "nfxid/modules/auth/domain/profile_occupation"
)

type Service struct {
	occupationRepo  *occupationDomain.Repo
	occupationQuery occupationDomain.Query
}

func NewService(
	occupationRepo *occupationDomain.Repo,
	occupationQuery occupationDomain.Query,
) *Service {
	return &Service{
		occupationRepo:  occupationRepo,
		occupationQuery: occupationQuery,
	}
}
