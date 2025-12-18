package user_permission

import (
	userPermissionDomain "nfxid/modules/permission/domain/user_permission"
)

type Service struct {
	userPermissionRepo  *userPermissionDomain.Repo
	userPermissionQuery *userPermissionDomain.Query
}

func NewService(
	userPermissionRepo *userPermissionDomain.Repo,
	userPermissionQuery *userPermissionDomain.Query,
) *Service {
	return &Service{
		userPermissionRepo:  userPermissionRepo,
		userPermissionQuery: userPermissionQuery,
	}
}
