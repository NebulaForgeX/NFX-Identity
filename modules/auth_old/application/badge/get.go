package badge

import (
	"context"
	badgeViews "nfxid/modules/auth/application/badge/views"
	badgeDomain "nfxid/modules/auth/domain/badge"

	"github.com/google/uuid"
)

func (s *Service) GetBadge(ctx context.Context, badgeID uuid.UUID) (badgeViews.BadgeView, error) {
	domainView, err := s.badgeQuery.Single.ByID(ctx, badgeID)
	if err != nil {
		return badgeViews.BadgeView{}, err
	}
	return badgeViews.BadgeViewMapper(*domainView), nil
}

func (s *Service) GetBadgeByName(ctx context.Context, name string) (badgeViews.BadgeView, error) {
	domainView, err := s.badgeQuery.Single.ByName(ctx, name)
	if err != nil {
		return badgeViews.BadgeView{}, err
	}
	return badgeViews.BadgeViewMapper(*domainView), nil
}

type GetBadgeListResult struct {
	Items []badgeViews.BadgeView
	Total int64
}

func (s *Service) GetBadgeList(ctx context.Context, q badgeDomain.ListQuery) (GetBadgeListResult, error) {
	q.Normalize()
	domainViews, total, err := s.badgeQuery.List.Generic(ctx, q)
	if err != nil {
		return GetBadgeListResult{}, err
	}
	items := make([]badgeViews.BadgeView, len(domainViews))
	for i, v := range domainViews {
		items[i] = badgeViews.BadgeViewMapper(*v)
	}
	return GetBadgeListResult{
		Items: items,
		Total: total,
	}, nil
}
