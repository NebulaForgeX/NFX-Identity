package queries

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	userRoleDomainViews "nfxid/modules/auth/domain/user_role/views"
	"nfxid/pkgs/query"

	"github.com/google/uuid"
)

type UserRoleQuery interface {
	GetByID(ctx context.Context, userRoleID uuid.UUID) (userRoleDomainViews.UserRoleView, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]userRoleDomainViews.UserRoleView, error)
	GetByRoleID(ctx context.Context, roleID uuid.UUID) ([]userRoleDomainViews.UserRoleView, error)
	GetList(ctx context.Context, q UserRoleListQuery) ([]userRoleDomainViews.UserRoleView, int64, error)
	GetCount(ctx context.Context) (int64, error)
}

type SortField int

const (
	_ SortField = iota
	SortByCreatedTime
	SortFieldMax
)

var allowedSortFields = query.MakeRangeSet(SortFieldMax)

const (
	defaultLimit = 20
	maxLimit     = 100
)

type UserRoleListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	UserRoleIDs []uuid.UUID `json:"user_role_ids,omitempty"`
	UserIDs     []uuid.UUID `json:"user_ids,omitempty"`
	RoleIDs     []uuid.UUID `json:"role_ids,omitempty"`
}

func (q *UserRoleListQuery) Normalize() {
	q.DomainPagination.Normalize(maxLimit, defaultLimit)
	q.DomainSorts.Normalize(
		allowedSortFields,
		&query.DomainSort[SortField]{Field: SortByCreatedTime, Order: "desc"},
	)
}

func (q UserRoleListQuery) CacheKey() string {
	b, _ := json.Marshal(q)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

