package reqdto

import (
	"strings"

	profileQuery "nfxidentity/modules/auth/query/profile"
	"nfxidentity/pkgs/query"

	"github.com/google/uuid"
)

// SearchProfiles is the request body for profile search endpoints
// (POST /auth/me/profiles/search and POST /auth/me/authority-profiles/search).
type SearchProfiles struct {
	Offset int     `json:"offset,omitempty"`
	Limit  int     `json:"limit,omitempty"`
	Query  *string `json:"query,omitempty"`
}

func (r *SearchProfiles) Sanitize() {
	if r.Query == nil {
		return
	}
	trimmed := strings.TrimSpace(*r.Query)
	if trimmed == "" {
		r.Query = nil
		return
	}
	r.Query = &trimmed
}

func (r SearchProfiles) ToQuery(excludeProfileID uuid.UUID) profileQuery.ListQuery {
	q := profileQuery.ListQuery{
		DomainPagination: query.DomainPagination{
			Offset: r.Offset,
			Limit:  r.Limit,
		},
		Search:            r.Query,
		ExcludeProfileIDs: []uuid.UUID{excludeProfileID},
	}
	q.Normalize()
	return q
}
