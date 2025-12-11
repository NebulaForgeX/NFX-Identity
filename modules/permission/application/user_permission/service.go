package user_permission

import (
	userPermissionQueries "nfxid/modules/permission/application/user_permission/queries"
	userPermissionDomain "nfxid/modules/permission/domain/user_permission"
)

type Service struct {
	userPermissionRepo  userPermissionDomain.Repo
	userPermissionQuery userPermissionQueries.UserPermissionQuery
}

func NewService(
	userPermissionRepo userPermissionDomain.Repo,
	userPermissionQuery userPermissionQueries.UserPermissionQuery,
) *Service {
	return &Service{
		userPermissionRepo:  userPermissionRepo,
		userPermissionQuery: userPermissionQuery,
	}
}

