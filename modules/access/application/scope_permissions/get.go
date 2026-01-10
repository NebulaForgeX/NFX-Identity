package scope_permissions

import (
	"context"
	scopePermissionResult "nfxid/modules/access/application/scope_permissions/results"

	"github.com/google/uuid"
)

// GetScopePermission 根据ID获取作用域权限
func (s *Service) GetScopePermission(ctx context.Context, scopePermissionID uuid.UUID) (scopePermissionResult.ScopePermissionRO, error) {
	domainEntity, err := s.scopePermissionRepo.Get.ByID(ctx, scopePermissionID)
	if err != nil {
		return scopePermissionResult.ScopePermissionRO{}, err
	}
	return scopePermissionResult.ScopePermissionMapper(domainEntity), nil
}

// GetScopePermissionByScopeAndPermission 根据作用域和权限ID获取作用域权限
func (s *Service) GetScopePermissionByScopeAndPermission(ctx context.Context, scope string, permissionID uuid.UUID) (scopePermissionResult.ScopePermissionRO, error) {
	domainEntity, err := s.scopePermissionRepo.Get.ByScopeAndPermissionID(ctx, scope, permissionID)
	if err != nil {
		return scopePermissionResult.ScopePermissionRO{}, err
	}
	return scopePermissionResult.ScopePermissionMapper(domainEntity), nil
}
