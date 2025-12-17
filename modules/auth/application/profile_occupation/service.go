package profile_occupation

import (
	occupationQueries "nfxid/modules/auth/application/profile_occupation/queries"
	occupationDomain "nfxid/modules/auth/domain/profile_occupation"
)

type Service struct {
	occupationRepo  *occupationDomain.Repo
	occupationQuery occupationQueries.OccupationQuery
}

func NewService(
	occupationRepo *occupationDomain.Repo,
	occupationQuery occupationQueries.OccupationQuery,
) *Service {
	return &Service{
		occupationRepo:  occupationRepo,
		occupationQuery: occupationQuery,
	}
}
