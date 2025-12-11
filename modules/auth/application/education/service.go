package education

import (
	"context"
	educationQueries "nfxid/modules/auth/application/education/queries"
	educationViews "nfxid/modules/auth/application/education/views"
	educationDomain "nfxid/modules/auth/domain/education"
	"nfxid/pkgs/cache"
	"nfxid/pkgs/eventbus"

	"github.com/google/uuid"
)

type Service struct {
	educationRepo  educationDomain.Repo
	educationQuery educationQueries.EducationQuery
	busPublisher   *eventbus.BusPublisher
	cacheSet       cache.CacheSet[educationViews.EducationView, uuid.UUID]
}

func NewService(
	educationRepo educationDomain.Repo,
	educationQuery educationQueries.EducationQuery,
	busPublisher *eventbus.BusPublisher,
	cacheSet cache.CacheSet[educationViews.EducationView, uuid.UUID],
) *Service {
	return &Service{
		educationRepo:  educationRepo,
		educationQuery: educationQuery,
		busPublisher:   busPublisher,
		cacheSet:       cacheSet,
	}
}

// InvalidateEducationCache 清除教育经历缓存
func (s *Service) InvalidateEducationCache(ctx context.Context, educationID uuid.UUID) error {
	return s.cacheSet.EntityCache().Invalidate(ctx, educationID)
}
