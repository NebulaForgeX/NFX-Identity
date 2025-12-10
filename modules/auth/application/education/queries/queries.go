package queries

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	educationDomainViews "nebulaid/modules/auth/domain/education/views"
	"nebulaid/pkgs/query"

	"github.com/google/uuid"
)

type EducationQuery interface {
	GetByID(ctx context.Context, educationID uuid.UUID) (educationDomainViews.EducationView, error)
	GetByProfileID(ctx context.Context, profileID uuid.UUID) ([]educationDomainViews.EducationView, error)
	GetList(ctx context.Context, q EducationListQuery) ([]educationDomainViews.EducationView, int64, error)
	GetCount(ctx context.Context) (int64, error)
}

type SortField int

const (
	_ SortField = iota
	SortByCreatedTime
	SortByStartDate
	SortBySchool
	SortFieldMax
)

var allowedSortFields = query.MakeRangeSet(SortFieldMax)

const (
	defaultLimit = 20
	maxLimit     = 100
)

type EducationListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	EducationIDs []uuid.UUID `json:"education_ids,omitempty"`
	ProfileIDs   []uuid.UUID `json:"profile_ids,omitempty"`
	Search       *string     `json:"search,omitempty"`
}

func (q *EducationListQuery) Normalize() {
	q.DomainPagination.Normalize(maxLimit, defaultLimit)
	q.DomainSorts.Normalize(
		allowedSortFields,
		&query.DomainSort[SortField]{Field: SortByStartDate, Order: "desc"},
	)
}

func (q EducationListQuery) CacheKey() string {
	b, _ := json.Marshal(q)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
