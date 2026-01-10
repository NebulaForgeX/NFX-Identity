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

// GetRolePermissionsByRoleID 根据角色ID获取角色权限列表
func (s *Service) GetRolePermissionsByRoleID(ctx context.Context, roleID uuid.UUID) ([]rolePermissionResult.RolePermissionRO, error) {
	domainEntities, err := s.rolePermissionRepo.Get.ByRoleID(ctx, roleID)
	if err != nil {
		return nil, err
	}

	results := make([]rolePermissionResult.RolePermissionRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = rolePermissionResult.RolePermissionMapper(entity)
	}
	return results, nil
}

// GetRolePermissionsByPermissionID 根据权限ID获取角色权限列表
func (s *Service) GetRolePermissionsByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]rolePermissionResult.RolePermissionRO, error) {
	domainEntities, err := s.rolePermissionRepo.Get.ByPermissionID(ctx, permissionID)
	if err != nil {
		return nil, err
	}

	results := make([]rolePermissionResult.RolePermissionRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = rolePermissionResult.RolePermissionMapper(entity)
	}
	return results, nil
}
