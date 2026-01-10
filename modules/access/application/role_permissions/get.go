package role_permissions

import (
	"context"
	rolePermissionResult "nfxid/modules/access/application/role_permissions/results"

	"github.com/google/uuid"
)

// GetRolePermission 根据ID获取角色权限
func (s *Service) GetRolePermission(ctx context.Context, rolePermissionID uuid.UUID) (rolePermissionResult.RolePermissionRO, error) {
	domainEntity, err := s.rolePermissionRepo.Get.ByID(ctx, rolePermissionID)
	if err != nil {
		return rolePermissionResult.RolePermissionRO{}, err
	}
	return rolePermissionResult.RolePermissionMapper(domainEntity), nil
}

// GetRolePermissionByRoleAndPermission 根据角色ID和权限ID获取角色权限
func (s *Service) GetRolePermissionByRoleAndPermission(ctx context.Context, roleID, permissionID uuid.UUID) (rolePermissionResult.RolePermissionRO, error) {
	domainEntity, err := s.rolePermissionRepo.Get.ByRoleIDAndPermissionID(ctx, roleID, permissionID)
	if err != nil {
		return rolePermissionResult.RolePermissionRO{}, err
	}
	return rolePermissionResult.RolePermissionMapper(domainEntity), nil
}
