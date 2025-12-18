package profile_badge

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"nfxid/modules/auth/domain/profile_badge/views"
	userDomainViews "nfxid/modules/auth/domain/user/views"
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
	ByID(ctx context.Context, profileBadgeID uuid.UUID) (*views.ProfileBadgeView, error)
	ByProfileID(ctx context.Context, profileID uuid.UUID) ([]*views.ProfileBadgeView, error)
	ByBadgeID(ctx context.Context, badgeID uuid.UUID) ([]*views.ProfileBadgeView, error)
	UserBadges(ctx context.Context, userID uuid.UUID) ([]*userDomainViews.UserBadgesView, error)
}

// List 定义列表查询相关的方法
type List interface {
	Generic(ctx context.Context, q ListQuery) ([]*views.ProfileBadgeView, int64, error)
}

// Count 定义计数相关的方法
type Count interface {
	All(ctx context.Context) (int64, error)
}

type SortField int

const (
	_ SortField = iota
	SortByCreatedTime
	SortByEarnedAt
	SortFieldMax
)

var allowedSortFields = query.MakeRangeSet(SortFieldMax)

const (
	defaultLimit = 20
	maxLimit     = 100
)

// ListQuery 用户徽章关联列表查询参数
type ListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	ProfileBadgeIDs []uuid.UUID `json:"profile_badge_ids,omitempty"`
	ProfileIDs      []uuid.UUID `json:"profile_ids,omitempty"`
	BadgeIDs        []uuid.UUID `json:"badge_ids,omitempty"`
}

func (q *ListQuery) Normalize() {
	q.DomainPagination.Normalize(maxLimit, defaultLimit)
	q.DomainSorts.Normalize(
		allowedSortFields,
		&query.DomainSort[SortField]{Field: SortByEarnedAt, Order: "desc"},
	)
}

func (q ListQuery) CacheKey() string {
	b, _ := json.Marshal(q)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
