package badge

import (
	"context"
	badgeViews "nfxid/modules/auth/application/badge/views"
	badgeDomain "nfxid/modules/auth/domain/badge"
	"nfxid/pkgs/cache"

	"github.com/google/uuid"
)

type Service struct {
	badgeRepo  *badgeDomain.Repo
	badgeQuery *badgeDomain.Query
	cacheSet   cache.CacheSet[badgeViews.BadgeView, uuid.UUID]
}

func NewService(
	badgeRepo *badgeDomain.Repo,
	badgeQuery *badgeDomain.Query,
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
