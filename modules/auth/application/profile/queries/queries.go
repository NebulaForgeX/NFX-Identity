package queries

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	profileDomainViews "nfxid/modules/auth/domain/profile/views"
	"nfxid/pkgs/query"

	"github.com/google/uuid"
)

type ProfileQuery interface {
	GetByID(ctx context.Context, profileID uuid.UUID) (profileDomainViews.ProfileView, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (profileDomainViews.ProfileView, error)
	GetList(ctx context.Context, q ProfileListQuery) ([]profileDomainViews.ProfileView, int64, error)
	GetCount(ctx context.Context) (int64, error)
}

type SortField int

const (
	_ SortField = iota
	SortByCreatedTime
	SortByDisplayName
	SortByNickname
	SortFieldMax
)

var allowedSortFields = query.MakeRangeSet(SortFieldMax)

const (
	defaultLimit = 20
	maxLimit     = 100
)

type ProfileListQuery struct {
	query.DomainPagination
	query.DomainSorts[SortField]

	ProfileIDs []uuid.UUID `json:"profile_ids,omitempty"`
	UserIDs    []uuid.UUID `json:"user_ids,omitempty"`
	Search     *string     `json:"search,omitempty"`
}

func (q *ProfileListQuery) Normalize() {
	q.DomainPagination.Normalize(maxLimit, defaultLimit)
	q.DomainSorts.Normalize(
		allowedSortFields,
		&query.DomainSort[SortField]{Field: SortByCreatedTime, Order: "desc"},
	)
}

func (q ProfileListQuery) CacheKey() string {
	b, _ := json.Marshal(q)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
