package permission

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"nfxid/modules/permission/domain/permission/views"
	"nfxid/pkgs/query"

	"github.com/google/uuid"
)

// Query 定义查询领域视图的结构体（CQRS Read Side）
type Query struct {
	Single Single
	List   List
}

// Single 定义单个查询相关的方法
type Single interface {
	ByID(ctx context.Context, permissionID uuid.UUID) (*views.PermissionView, error)
	ByTag(ctx context.Context, tag string) (*views.PermissionView, error)
}

// List 定义列表查询相关的方法
type List interface {
	Generic(ctx context.Context, q ListQuery) ([]*views.PermissionView, int64, error)
	ByTags(ctx context.Context, tags []string) ([]*views.PermissionView, error)
	ByCategory(ctx context.Context, category string) ([]*views.PermissionView, error)
}

type SortField int

const (
	_ SortField = iota
	SortByCreatedTime
	SortByTag
	SortByCategory
	SortFieldMax
)

var allowedSortFields = query.MakeRangeSet(SortFieldMax)

const (
	defaultLimit = 20
	maxLimit     = 100
)

// ListQuery 权限列表查询参数
type ListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	PermissionIDs []uuid.UUID `json:"permission_ids,omitempty"`
	Category      *string     `json:"category,omitempty"`
	IsSystem      *bool       `json:"is_system,omitempty"`
	Search        *string     `json:"search,omitempty"`
}

func (q *ListQuery) Normalize() {
	q.DomainPagination.Normalize(maxLimit, defaultLimit)
	q.DomainSorts.Normalize(
		allowedSortFields,
		&query.DomainSort[SortField]{Field: SortByCategory, Order: "asc"},
	)
}

func (q ListQuery) CacheKey() string {
	b, _ := json.Marshal(q)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
