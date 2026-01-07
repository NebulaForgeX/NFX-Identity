package education

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"nfxid/modules/auth/domain/profile_education/views"
	"nfxid/pkgs/query"

	"github.com/google/uuid"
)

// Query 定义查询领域视图的结构体（CQRS Read Side）
type Query struct {
	Single Single
	List   List
	Count  Count
}

// Single 定义单个查询相关的方法
type Single interface {
	ByID(ctx context.Context, educationID uuid.UUID) (*views.EducationView, error)
}

// List 定义列表查询相关的方法
type List interface {
	Generic(ctx context.Context, q ListQuery) ([]*views.EducationView, int64, error)
	ByProfileID(ctx context.Context, profileID uuid.UUID) ([]*views.EducationView, error)
}

// Count 定义计数相关的方法
type Count interface {
	All(ctx context.Context) (int64, error)
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

// ListQuery 教育经历列表查询参数
type ListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	EducationIDs []uuid.UUID `json:"education_ids,omitempty"`
	ProfileIDs   []uuid.UUID `json:"profile_ids,omitempty"`
	Search       *string     `json:"search,omitempty"`
}

func (q *ListQuery) Normalize() {
	q.DomainPagination.Normalize(maxLimit, defaultLimit)
	q.DomainSorts.Normalize(
		allowedSortFields,
		&query.DomainSort[SortField]{Field: SortByStartDate, Order: "desc"},
	)
}

func (q ListQuery) CacheKey() string {
	b, _ := json.Marshal(q)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
