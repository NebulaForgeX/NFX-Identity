package occupation

import (
	occupationQueries "nebulaid/modules/auth/application/occupation/queries"
	occupationDomain "nebulaid/modules/auth/domain/occupation"
)

type Service struct {
	occupationRepo  occupationDomain.Repo
	occupationQuery occupationQueries.OccupationQuery
}

func NewService(
	occupationRepo occupationDomain.Repo,
	occupationQuery occupationQueries.OccupationQuery,
) *Service {
	return &Service{
		occupationRepo:  occupationRepo,
		occupationQuery: occupationQuery,
	}
}
