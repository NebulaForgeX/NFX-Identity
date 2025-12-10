package queries

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	badgeDomainViews "nebulaid/modules/auth/domain/badge/views"
	"nebulaid/pkgs/query"

	"github.com/google/uuid"
)

type BadgeQuery interface {
	GetByID(ctx context.Context, badgeID uuid.UUID) (badgeDomainViews.BadgeView, error)
	GetByName(ctx context.Context, name string) (badgeDomainViews.BadgeView, error)
	GetList(ctx context.Context, q BadgeListQuery) ([]badgeDomainViews.BadgeView, int64, error)
}

type SortField int

const (
	_ SortField = iota
	SortByCreatedTime
	SortByName
	SortFieldMax
)

var allowedSortFields = query.MakeRangeSet(SortFieldMax)

const (
	defaultLimit = 20
	maxLimit     = 100
)

type BadgeListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	BadgeIDs []uuid.UUID `json:"badge_ids,omitempty"`
	Search   *string     `json:"search,omitempty"`
	Category *string     `json:"category,omitempty"`
	IsSystem *bool       `json:"is_system,omitempty"`
}

func (q *BadgeListQuery) Normalize() {
	q.DomainPagination.Normalize(maxLimit, defaultLimit)
	q.DomainSorts.Normalize(
		allowedSortFields,
		&query.DomainSort[SortField]{Field: SortByCreatedTime, Order: "desc"},
	)
}

func (q BadgeListQuery) CacheKey() string {
	b, _ := json.Marshal(q)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
