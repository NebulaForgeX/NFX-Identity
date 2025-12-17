package user_role

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"nfxid/modules/auth/domain/user_role/views"
	"nfxid/pkgs/query"

	"github.com/google/uuid"
)

// Query 定义查询领域视图的接口（CQRS Read Side）
type Query interface {
	ByID(ctx context.Context, userRoleID uuid.UUID) (views.UserRoleView, error)
	ByUserID(ctx context.Context, userID uuid.UUID) ([]views.UserRoleView, error)
	ByRoleID(ctx context.Context, roleID uuid.UUID) ([]views.UserRoleView, error)
	List(ctx context.Context, q ListQuery) ([]views.UserRoleView, int64, error)
	Count(ctx context.Context) (int64, error)
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

// ListQuery 用户角色关联列表查询参数
type ListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	UserRoleIDs []uuid.UUID `json:"user_role_ids,omitempty"`
	UserIDs     []uuid.UUID `json:"user_ids,omitempty"`
	RoleIDs     []uuid.UUID `json:"role_ids,omitempty"`
}

func (q *ListQuery) Normalize() {
	q.DomainPagination.Normalize(maxLimit, defaultLimit)
	q.DomainSorts.Normalize(
		allowedSortFields,
		&query.DomainSort[SortField]{Field: SortByCreatedTime, Order: "desc"},
	)
}

func (q ListQuery) CacheKey() string {
	b, _ := json.Marshal(q)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
