package user

import (
	"context"

	"nebulaid/modules/auth/application/user/queries"
	"nebulaid/modules/auth/application/user/views"

	"github.com/google/uuid"
)

func (s *Service) GetUser(ctx context.Context, userID uuid.UUID) (views.UserView, error) {
	domainView, err := s.userQuery.GetByID(ctx, userID)
	if err != nil {
		return views.UserView{}, err
	}
	return views.UserViewMapper(domainView), nil
}

type GetUserListResult struct {
	Items []views.UserView
	Total int64
}

func (s *Service) GetUserList(ctx context.Context, q queries.UserListQuery) (GetUserListResult, error) {
	q.Normalize()
	domainViews, total, err := s.userQuery.GetList(ctx, q)
	if err != nil {
		return GetUserListResult{}, err
	}
	items := make([]views.UserView, len(domainViews))
	for i, v := range domainViews {
		items[i] = views.UserViewMapper(v)
	}
	return GetUserListResult{
		Items: items,
		Total: total,
	}, nil
}
