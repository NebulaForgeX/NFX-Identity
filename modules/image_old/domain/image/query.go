package image

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"nfxid/modules/image/domain/image/views"
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
	ByID(ctx context.Context, imageID uuid.UUID) (*views.ImageView, error)
}

// List 定义列表查询相关的方法
type List interface {
	Generic(ctx context.Context, q ListQuery) ([]*views.ImageView, int64, error)
}

type SortField int

const (
	_ SortField = iota
	SortByCreatedTime
	SortByUpdatedAt
	SortBySize
	SortFieldMax
)

var allowedSortFields = query.MakeRangeSet(SortFieldMax)

const (
	defaultLimit = 20
	maxLimit     = 100
)

// ListQuery 图片列表查询参数
type ListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	ImageIDs     []uuid.UUID `json:"image_ids,omitempty"`
	TypeID       *uuid.UUID  `json:"type_id,omitempty"`
	UserID       *uuid.UUID  `json:"user_id,omitempty"`
	SourceDomain *string     `json:"source_domain,omitempty"`
	IsPublic     *bool       `json:"is_public,omitempty"`
	Search       *string     `json:"search,omitempty"`
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
