package user_role

import (
	"context"
	userRoleViews "nfxid/modules/auth/application/user_role/views"
	userRoleDomain "nfxid/modules/auth/domain/user_role"

	"github.com/google/uuid"
)

func (s *Service) GetUserRole(ctx context.Context, userRoleID uuid.UUID) (userRoleViews.UserRoleView, error) {
	domainView, err := s.userRoleQuery.ByID(ctx, userRoleID)
	if err != nil {
		return userRoleViews.UserRoleView{}, err
	}
	return userRoleViews.UserRoleViewMapper(domainView), nil
}

func (s *Service) GetUserRolesByUserID(ctx context.Context, userID uuid.UUID) ([]userRoleViews.UserRoleView, error) {
	domainViews, err := s.userRoleQuery.ByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	result := make([]userRoleViews.UserRoleView, len(domainViews))
	for i, v := range domainViews {
		result[i] = userRoleViews.UserRoleViewMapper(v)
	}
	return result, nil
}

func (s *Service) GetUserRolesByRoleID(ctx context.Context, roleID uuid.UUID) ([]userRoleViews.UserRoleView, error) {
	domainViews, err := s.userRoleQuery.ByRoleID(ctx, roleID)
	if err != nil {
		return nil, err
	}
	result := make([]userRoleViews.UserRoleView, len(domainViews))
	for i, v := range domainViews {
		result[i] = userRoleViews.UserRoleViewMapper(v)
	}
	return result, nil
}

type GetUserRoleListResult struct {
	Items []userRoleViews.UserRoleView
	Total int64
}

func (s *Service) GetUserRoleList(ctx context.Context, q userRoleDomain.ListQuery) (GetUserRoleListResult, error) {
	q.Normalize()
	domainViews, total, err := s.userRoleQuery.List(ctx, q)
	if err != nil {
		return GetUserRoleListResult{}, err
	}
	items := make([]userRoleViews.UserRoleView, len(domainViews))
	for i, v := range domainViews {
		items[i] = userRoleViews.UserRoleViewMapper(v)
	}
	return GetUserRoleListResult{
		Items: items,
		Total: total,
	}, nil
}
