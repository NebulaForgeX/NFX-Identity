package profile_badge

import (
	"context"
	profileBadgeQueries "nebulaid/modules/auth/application/profile_badge/queries"
	profileBadgeViews "nebulaid/modules/auth/application/profile_badge/views"

	"github.com/google/uuid"
)

func (s *Service) GetProfileBadge(ctx context.Context, profileBadgeID uuid.UUID) (profileBadgeViews.ProfileBadgeView, error) {
	domainView, err := s.profileBadgeQuery.GetByID(ctx, profileBadgeID)
	if err != nil {
		return profileBadgeViews.ProfileBadgeView{}, err
	}
	return profileBadgeViews.ProfileBadgeViewMapper(domainView), nil
}

func (s *Service) GetProfileBadgesByProfileID(ctx context.Context, profileID uuid.UUID) ([]profileBadgeViews.ProfileBadgeView, error) {
	domainViews, err := s.profileBadgeQuery.GetByProfileID(ctx, profileID)
	if err != nil {
		return nil, err
	}
	result := make([]profileBadgeViews.ProfileBadgeView, len(domainViews))
	for i, v := range domainViews {
		result[i] = profileBadgeViews.ProfileBadgeViewMapper(v)
	}
	return result, nil
}

func (s *Service) GetProfileBadgesByBadgeID(ctx context.Context, badgeID uuid.UUID) ([]profileBadgeViews.ProfileBadgeView, error) {
	domainViews, err := s.profileBadgeQuery.GetByBadgeID(ctx, badgeID)
	if err != nil {
		return nil, err
	}
	result := make([]profileBadgeViews.ProfileBadgeView, len(domainViews))
	for i, v := range domainViews {
		result[i] = profileBadgeViews.ProfileBadgeViewMapper(v)
	}
	return result, nil
}

func (s *Service) GetUserBadges(ctx context.Context, userID uuid.UUID) ([]profileBadgeViews.UserBadgeView, error) {
	domainViews, err := s.profileBadgeQuery.GetUserBadges(ctx, userID)
	if err != nil {
		return nil, err
	}
	result := make([]profileBadgeViews.UserBadgeView, len(domainViews))
	for i, v := range domainViews {
		result[i] = profileBadgeViews.UserBadgeViewMapper(v)
	}
	return result, nil
}

type GetProfileBadgeListResult struct {
	Items []profileBadgeViews.ProfileBadgeView
	Total int64
}

func (s *Service) GetProfileBadgeList(ctx context.Context, q profileBadgeQueries.ProfileBadgeListQuery) (GetProfileBadgeListResult, error) {
	q.Normalize()
	domainViews, total, err := s.profileBadgeQuery.GetList(ctx, q)
	if err != nil {
		return GetProfileBadgeListResult{}, err
	}
	items := make([]profileBadgeViews.ProfileBadgeView, len(domainViews))
	for i, v := range domainViews {
		items[i] = profileBadgeViews.ProfileBadgeViewMapper(v)
	}
	return GetProfileBadgeListResult{
		Items: items,
		Total: total,
	}, nil
}
