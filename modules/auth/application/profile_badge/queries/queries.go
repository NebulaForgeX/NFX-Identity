package queries

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	profileBadgeDomainViews "nebulaid/modules/auth/domain/profile_badge/views"
	userDomainViews "nebulaid/modules/auth/domain/user/views"
	"nebulaid/pkgs/query"

	"github.com/google/uuid"
)

type ProfileBadgeQuery interface {
	GetByID(ctx context.Context, profileBadgeID uuid.UUID) (profileBadgeDomainViews.ProfileBadgeView, error)
	GetByProfileID(ctx context.Context, profileID uuid.UUID) ([]profileBadgeDomainViews.ProfileBadgeView, error)
	GetByBadgeID(ctx context.Context, badgeID uuid.UUID) ([]profileBadgeDomainViews.ProfileBadgeView, error)
	GetUserBadges(ctx context.Context, userID uuid.UUID) ([]userDomainViews.UserBadgesView, error)
	GetList(ctx context.Context, q ProfileBadgeListQuery) ([]profileBadgeDomainViews.ProfileBadgeView, int64, error)
	GetCount(ctx context.Context) (int64, error)
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

type ProfileBadgeListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	ProfileBadgeIDs []uuid.UUID `json:"profile_badge_ids,omitempty"`
	ProfileIDs      []uuid.UUID `json:"profile_ids,omitempty"`
	BadgeIDs        []uuid.UUID `json:"badge_ids,omitempty"`
}

func (q *ProfileBadgeListQuery) Normalize() {
	q.DomainPagination.Normalize(maxLimit, defaultLimit)
	q.DomainSorts.Normalize(
		allowedSortFields,
		&query.DomainSort[SortField]{Field: SortByEarnedAt, Order: "desc"},
	)
}

func (q ProfileBadgeListQuery) CacheKey() string {
	b, _ := json.Marshal(q)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
