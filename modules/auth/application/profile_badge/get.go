package profile_badge

import (
	"context"
	profileBadgeViews "nfxid/modules/auth/application/profile_badge/views"
	profileBadgeDomain "nfxid/modules/auth/domain/profile_badge"

	"github.com/google/uuid"
)

func (s *Service) GetProfileBadge(ctx context.Context, profileBadgeID uuid.UUID) (profileBadgeViews.ProfileBadgeView, error) {
	domainView, err := s.profileBadgeQuery.Single.ByID(ctx, profileBadgeID)
	if err != nil {
		return profileBadgeViews.ProfileBadgeView{}, err
	}
	return profileBadgeViews.ProfileBadgeViewMapper(*domainView), nil
}

func (s *Service) GetProfileBadgesByProfileID(ctx context.Context, profileID uuid.UUID) ([]profileBadgeViews.ProfileBadgeView, error) {
	domainViews, err := s.profileBadgeQuery.List.ByProfileID(ctx, profileID)
	if err != nil {
		return nil, err
	}
	result := make([]profileBadgeViews.ProfileBadgeView, len(domainViews))
	for i, v := range domainViews {
		result[i] = profileBadgeViews.ProfileBadgeViewMapper(*v)
	}
	return result, nil
}

func (s *Service) GetProfileBadgesByBadgeID(ctx context.Context, badgeID uuid.UUID) ([]profileBadgeViews.ProfileBadgeView, error) {
	domainViews, err := s.profileBadgeQuery.List.ByBadgeID(ctx, badgeID)
	if err != nil {
		return nil, err
	}
	result := make([]profileBadgeViews.ProfileBadgeView, len(domainViews))
	for i, v := range domainViews {
		result[i] = profileBadgeViews.ProfileBadgeViewMapper(*v)
	}
	return result, nil
}

func (s *Service) GetUserBadges(ctx context.Context, userID uuid.UUID) ([]profileBadgeViews.UserBadgeView, error) {
	domainViews, err := s.profileBadgeQuery.List.UserBadges(ctx, userID)
	if err != nil {
		return nil, err
	}
	result := make([]profileBadgeViews.UserBadgeView, len(domainViews))
	for i, v := range domainViews {
		result[i] = profileBadgeViews.UserBadgeViewMapper(*v)
	}
	return result, nil
}

type GetProfileBadgeListResult struct {
	Items []profileBadgeViews.ProfileBadgeView
	Total int64
}

func (s *Service) GetProfileBadgeList(ctx context.Context, q profileBadgeDomain.ListQuery) (GetProfileBadgeListResult, error) {
	q.Normalize()
	domainViews, total, err := s.profileBadgeQuery.List.Generic(ctx, q)
	if err != nil {
		return GetProfileBadgeListResult{}, err
	}
	items := make([]profileBadgeViews.ProfileBadgeView, len(domainViews))
	for i, v := range domainViews {
		items[i] = profileBadgeViews.ProfileBadgeViewMapper(*v)
	}
	return GetProfileBadgeListResult{
		Items: items,
		Total: total,
	}, nil
}
