package role

import (
	"context"
	roleViews "nfxid/modules/auth/application/role/views"
	roleDomain "nfxid/modules/auth/domain/role"

	"github.com/google/uuid"
)

func (s *Service) GetRole(ctx context.Context, roleID uuid.UUID) (roleViews.RoleView, error) {
	domainView, err := s.roleQuery.ByID(ctx, roleID)
	if err != nil {
		return roleViews.RoleView{}, err
	}
	return roleViews.RoleViewMapper(domainView), nil
}

func (s *Service) GetRoleByName(ctx context.Context, name string) (roleViews.RoleView, error) {
	domainView, err := s.roleQuery.ByName(ctx, name)
	if err != nil {
		return roleViews.RoleView{}, err
	}
	return roleViews.RoleViewMapper(domainView), nil
}

type GetRoleListResult struct {
	Items []roleViews.RoleView
	Total int64
}

func (s *Service) GetRoleList(ctx context.Context, q roleDomain.ListQuery) (GetRoleListResult, error) {
	q.Normalize()
	domainViews, total, err := s.roleQuery.List(ctx, q)
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
