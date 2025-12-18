package profile_education

import (
	"context"
	educationViews "nfxid/modules/auth/application/profile_education/views"
	educationDomain "nfxid/modules/auth/domain/profile_education"
	"nfxid/pkgs/cache"
	"nfxid/pkgs/eventbus"

	"github.com/google/uuid"
)

type Service struct {
	educationRepo  *educationDomain.Repo
	educationQuery *educationDomain.Query
	busPublisher   *eventbus.BusPublisher
	cacheSet       cache.CacheSet[educationViews.EducationView, uuid.UUID]
}

func NewService(
	educationRepo *educationDomain.Repo,
	educationQuery *educationDomain.Query,
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
