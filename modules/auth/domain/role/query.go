package role

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"nfxid/modules/auth/domain/role/views"
	"nfxid/pkgs/query"

	"github.com/google/uuid"
)

// Query 定义查询领域视图的接口（CQRS Read Side）
type Query interface {
	ByID(ctx context.Context, roleID uuid.UUID) (views.RoleView, error)
	ByName(ctx context.Context, name string) (views.RoleView, error)
	List(ctx context.Context, q ListQuery) ([]views.RoleView, int64, error)
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

// ListQuery 角色列表查询参数
type ListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	RoleIDs  []uuid.UUID `json:"role_ids,omitempty"`
	Search   *string     `json:"search,omitempty"`
	IsSystem *bool       `json:"is_system,omitempty"`
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
