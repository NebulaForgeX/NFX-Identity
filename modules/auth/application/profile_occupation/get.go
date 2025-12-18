package profile_occupation

import (
	"context"
	occupationViews "nfxid/modules/auth/application/profile_occupation/views"
	occupationDomain "nfxid/modules/auth/domain/profile_occupation"

	"github.com/google/uuid"
)

func (s *Service) GetOccupation(ctx context.Context, occupationID uuid.UUID) (occupationViews.OccupationView, error) {
	domainView, err := s.occupationQuery.Single.ByID(ctx, occupationID)
	if err != nil {
		return occupationViews.OccupationView{}, err
	}
	return occupationViews.OccupationViewMapper(*domainView), nil
}

func (s *Service) GetOccupationsByProfileID(ctx context.Context, profileID uuid.UUID) ([]occupationViews.OccupationView, error) {
	domainViews, err := s.occupationQuery.List.ByProfileID(ctx, profileID)
	if err != nil {
		return nil, err
	}
	result := make([]occupationViews.OccupationView, len(domainViews))
	for i, v := range domainViews {
		result[i] = occupationViews.OccupationViewMapper(*v)
	}
	return result, nil
}

type GetOccupationListResult struct {
	Items []occupationViews.OccupationView
	Total int64
}

func (s *Service) GetOccupationList(ctx context.Context, q occupationDomain.ListQuery) (GetOccupationListResult, error) {
	q.Normalize()
	domainViews, total, err := s.occupationQuery.List.Generic(ctx, q)
	if err != nil {
		return GetOccupationListResult{}, err
	}
	items := make([]occupationViews.OccupationView, len(domainViews))
	for i, v := range domainViews {
		items[i] = occupationViews.OccupationViewMapper(*v)
	}
	return GetOccupationListResult{
		Items: items,
		Total: total,
	}, nil
}
