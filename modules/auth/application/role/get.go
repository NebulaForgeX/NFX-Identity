package role

import (
	"context"
	roleQueries "nebulaid/modules/auth/application/role/queries"
	roleViews "nebulaid/modules/auth/application/role/views"

	"github.com/google/uuid"
)

func (s *Service) GetRole(ctx context.Context, roleID uuid.UUID) (roleViews.RoleView, error) {
	domainView, err := s.roleQuery.GetByID(ctx, roleID)
	if err != nil {
		return roleViews.RoleView{}, err
	}
	return roleViews.RoleViewMapper(domainView), nil
}

func (s *Service) GetRoleByName(ctx context.Context, name string) (roleViews.RoleView, error) {
	domainView, err := s.roleQuery.GetByName(ctx, name)
	if err != nil {
		return roleViews.RoleView{}, err
	}
	return roleViews.RoleViewMapper(domainView), nil
}

type GetRoleListResult struct {
	Items []roleViews.RoleView
	Total int64
}

func (s *Service) GetRoleList(ctx context.Context, q roleQueries.RoleListQuery) (GetRoleListResult, error) {
	q.Normalize()
	domainViews, total, err := s.roleQuery.GetList(ctx, q)
	if err != nil {
		return GetRoleListResult{}, err
	}
	items := make([]roleViews.RoleView, len(domainViews))
	for i, v := range domainViews {
		items[i] = roleViews.RoleViewMapper(v)
	}
	return GetRoleListResult{
		Items: items,
		Total: total,
	}, nil
}
