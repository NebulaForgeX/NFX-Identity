package permissions

import (
	"context"
	permissionResult "nfxid/modules/access/application/permissions/results"

	"github.com/google/uuid"
)

// GetPermission 根据ID获取权限
func (s *Service) GetPermission(ctx context.Context, permissionID uuid.UUID) (permissionResult.PermissionRO, error) {
	domainEntity, err := s.permissionRepo.Get.ByID(ctx, permissionID)
	if err != nil {
		return permissionResult.PermissionRO{}, err
	}
	return permissionResult.PermissionMapper(domainEntity), nil
}

// GetPermissionByKey 根据Key获取权限
func (s *Service) GetPermissionByKey(ctx context.Context, key string) (permissionResult.PermissionRO, error) {
	domainEntity, err := s.permissionRepo.Get.ByKey(ctx, key)
	if err != nil {
		return permissionResult.PermissionRO{}, err
	}
	return permissionResult.PermissionMapper(domainEntity), nil
}
