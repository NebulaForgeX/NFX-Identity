package queries

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	occupationDomainViews "nebulaid/modules/auth/domain/occupation/views"
	"nebulaid/pkgs/query"

	"github.com/google/uuid"
)

type OccupationQuery interface {
	GetByID(ctx context.Context, occupationID uuid.UUID) (occupationDomainViews.OccupationView, error)
	GetByProfileID(ctx context.Context, profileID uuid.UUID) ([]occupationDomainViews.OccupationView, error)
	GetList(ctx context.Context, q OccupationListQuery) ([]occupationDomainViews.OccupationView, int64, error)
	GetCount(ctx context.Context) (int64, error)
}

type SortField int

const (
	_ SortField = iota
	SortByCreatedTime
	SortByStartDate
	SortByCompany
	SortFieldMax
)

var allowedSortFields = query.MakeRangeSet(SortFieldMax)

const (
	defaultLimit = 20
	maxLimit     = 100
)

type OccupationListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	OccupationIDs []uuid.UUID `json:"occupation_ids,omitempty"`
	ProfileIDs    []uuid.UUID `json:"profile_ids,omitempty"`
	Search        *string     `json:"search,omitempty"`
}

func (q *OccupationListQuery) Normalize() {
	q.DomainPagination.Normalize(maxLimit, defaultLimit)
	q.DomainSorts.Normalize(
		allowedSortFields,
		&query.DomainSort[SortField]{Field: SortByStartDate, Order: "desc"},
	)
}

func (q OccupationListQuery) CacheKey() string {
	b, _ := json.Marshal(q)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
