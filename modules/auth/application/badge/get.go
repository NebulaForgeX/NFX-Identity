package badge

import (
	"context"
	badgeQueries "nebulaid/modules/auth/application/badge/queries"
	badgeViews "nebulaid/modules/auth/application/badge/views"

	"github.com/google/uuid"
)

func (s *Service) GetBadge(ctx context.Context, badgeID uuid.UUID) (badgeViews.BadgeView, error) {
	domainView, err := s.badgeQuery.GetByID(ctx, badgeID)
	if err != nil {
		return badgeViews.BadgeView{}, err
	}
	return badgeViews.BadgeViewMapper(domainView), nil
}

func (s *Service) GetBadgeByName(ctx context.Context, name string) (badgeViews.BadgeView, error) {
	domainView, err := s.badgeQuery.GetByName(ctx, name)
	if err != nil {
		return badgeViews.BadgeView{}, err
	}
	return badgeViews.BadgeViewMapper(domainView), nil
}

type GetBadgeListResult struct {
	Items []badgeViews.BadgeView
	Total int64
}

func (s *Service) GetBadgeList(ctx context.Context, q badgeQueries.BadgeListQuery) (GetBadgeListResult, error) {
	q.Normalize()
	domainViews, total, err := s.badgeQuery.GetList(ctx, q)
	if err != nil {
		return GetBadgeListResult{}, err
	}
	items := make([]badgeViews.BadgeView, len(domainViews))
	for i, v := range domainViews {
		items[i] = badgeViews.BadgeViewMapper(v)
	}
	return GetBadgeListResult{
		Items: items,
		Total: total,
	}, nil
}
