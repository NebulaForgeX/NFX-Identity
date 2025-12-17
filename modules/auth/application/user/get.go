package user

import (
	"context"
	"nfxid/modules/auth/application/user/views"
	userDomain "nfxid/modules/auth/domain/user"

	"github.com/google/uuid"
)

func (s *Service) GetUser(ctx context.Context, userID uuid.UUID) (views.UserView, error) {
	domainView, err := s.userQuery.ByID(ctx, userID)
	if err != nil {
		return views.UserView{}, err
	}
	return views.UserViewMapper(domainView), nil
}

type GetUserListResult struct {
	Items []views.UserView
	Total int64
}

func (s *Service) GetUserList(ctx context.Context, q userDomain.ListQuery) (GetUserListResult, error) {
	q.Normalize()
	domainViews, total, err := s.userQuery.List(ctx, q)
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
