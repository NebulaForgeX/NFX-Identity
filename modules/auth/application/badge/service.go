package badge

import (
	"context"
	badgeQueries "nebulaid/modules/auth/application/badge/queries"
	badgeViews "nebulaid/modules/auth/application/badge/views"
	"nebulaid/modules/auth/domain/badge"
	"nebulaid/pkgs/cache"

	"github.com/google/uuid"
)

type Service struct {
	badgeRepo  badge.Repo
	badgeQuery badgeQueries.BadgeQuery
	cacheSet   cache.CacheSet[badgeViews.BadgeView, uuid.UUID]
}

func NewService(
	badgeRepo badge.Repo,
	badgeQuery badgeQueries.BadgeQuery,
	cacheSet cache.CacheSet[badgeViews.BadgeView, uuid.UUID],
) *Service {
	return &Service{
		badgeRepo:  badgeRepo,
		badgeQuery: badgeQuery,
		cacheSet:   cacheSet,
	}
}

// InvalidateBadgeCache 清除徽章缓存
func (s *Service) InvalidateBadgeCache(ctx context.Context, badgeID uuid.UUID) error {
	return s.cacheSet.EntityCache().Invalidate(ctx, badgeID)
}
