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

// Query 定义查询领域视图的接口（CQRS Read Side）
type Query interface {
	ByID(ctx context.Context, educationID uuid.UUID) (views.EducationView, error)
	ByProfileID(ctx context.Context, profileID uuid.UUID) ([]views.EducationView, error)
	List(ctx context.Context, q ListQuery) ([]views.EducationView, int64, error)
	Count(ctx context.Context) (int64, error)
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
