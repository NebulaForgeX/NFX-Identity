package profile

import (
	"nfxidentity/pkgs/query"

	"github.com/google/uuid"
)

const (
	searchDefaultLimit = 20
	searchMaxLimit     = 50
)

type ListQuery struct {
	query.DomainPagination
	Search            *string     `json:"search,omitempty"`
	ExcludeProfileIDs []uuid.UUID `json:"exclude_profile_ids,omitempty"`
}

func (q *ListQuery) Normalize() {
	q.DomainPagination.Normalize(searchMaxLimit, searchDefaultLimit)
}
