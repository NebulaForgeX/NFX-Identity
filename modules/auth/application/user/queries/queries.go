package queries

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	userDomainViews "nebulaid/modules/auth/domain/user/views"
	"nebulaid/pkgs/query"

	"github.com/google/uuid"
)

type UserQuery interface {
	GetByID(ctx context.Context, userID uuid.UUID) (userDomainViews.UserView, error)
	GetByUsername(ctx context.Context, username string) (userDomainViews.UserView, error)
	GetByEmail(ctx context.Context, email string) (userDomainViews.UserView, error)
	GetList(ctx context.Context, q UserListQuery) ([]userDomainViews.UserView, int64, error)
	GetCount(ctx context.Context) (int64, error)
}

type SortField int

const (
	_ SortField = iota
	SortByCreatedTime
	SortByUsername
	SortByEmail
	SortFieldMax
)

var allowedSortFields = query.MakeRangeSet(SortFieldMax)

const (
	defaultLimit = 20
	maxLimit     = 100
)

type UserListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	UserIDs    []uuid.UUID `json:"user_ids,omitempty"`
	RoleIDs    []uuid.UUID `json:"role_ids,omitempty"`
	Status     []string    `json:"status,omitempty"`
	Search     *string     `json:"search,omitempty"`
	IsVerified *bool       `json:"is_verified,omitempty"`
}

func (q *UserListQuery) Normalize() {
	q.DomainPagination.Normalize(maxLimit, defaultLimit)
	q.DomainSorts.Normalize(
		allowedSortFields,
		&query.DomainSort[SortField]{Field: SortByCreatedTime, Order: "desc"},
	)
}

func (q UserListQuery) CacheKey() string {
	b, _ := json.Marshal(q)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
