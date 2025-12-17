package user

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	views "nfxid/modules/auth/domain/user/views"
	"nfxid/pkgs/query"

	"github.com/google/uuid"
)

// Query 定义查询领域视图的接口（CQRS Read Side）
type Query interface {
	ByID(ctx context.Context, userID uuid.UUID) (views.UserView, error)
	ByUsername(ctx context.Context, username string) (views.UserView, error)
	ByEmail(ctx context.Context, email string) (views.UserView, error)
	List(ctx context.Context, q ListQuery) ([]views.UserView, int64, error)
	Count(ctx context.Context) (int64, error)
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

// ListQuery 用户列表查询参数
type ListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	UserIDs    []uuid.UUID `json:"user_ids,omitempty"`
	RoleIDs    []uuid.UUID `json:"role_ids,omitempty"`
	Status     []string    `json:"status,omitempty"`
	Search     *string     `json:"search,omitempty"`
	IsVerified *bool       `json:"is_verified,omitempty"`
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
