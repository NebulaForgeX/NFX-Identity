package queries

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	roleDomainViews "nebulaid/modules/auth/domain/role/views"
	"nebulaid/pkgs/query"

	"github.com/google/uuid"
)

type RoleQuery interface {
	GetByID(ctx context.Context, roleID uuid.UUID) (roleDomainViews.RoleView, error)
	GetByName(ctx context.Context, name string) (roleDomainViews.RoleView, error)
	GetList(ctx context.Context, q RoleListQuery) ([]roleDomainViews.RoleView, int64, error)
	GetCount(ctx context.Context) (int64, error)
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

type RoleListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	RoleIDs  []uuid.UUID `json:"role_ids,omitempty"`
	Search   *string     `json:"search,omitempty"`
	IsSystem *bool       `json:"is_system,omitempty"`
}

func (q *RoleListQuery) Normalize() {
	q.DomainPagination.Normalize(maxLimit, defaultLimit)
	q.DomainSorts.Normalize(
		allowedSortFields,
		&query.DomainSort[SortField]{Field: SortByCreatedTime, Order: "desc"},
	)
}

func (q RoleListQuery) CacheKey() string {
	b, _ := json.Marshal(q)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
