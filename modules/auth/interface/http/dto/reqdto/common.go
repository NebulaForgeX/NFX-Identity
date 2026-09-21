package reqdto

import "nfxidentity/pkgs/query"

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func (p Pagination) domainPagination() query.DomainPagination {
	return query.DomainPagination{
		Limit:  p.Limit,
		Offset: p.Offset,
	}
}
