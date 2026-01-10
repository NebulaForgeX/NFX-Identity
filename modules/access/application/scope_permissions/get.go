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

// GetScopePermissionsByScope 根据作用域获取作用域权限列表
func (s *Service) GetScopePermissionsByScope(ctx context.Context, scope string) ([]scopePermissionResult.ScopePermissionRO, error) {
	domainEntities, err := s.scopePermissionRepo.Get.ByScope(ctx, scope)
	if err != nil {
		return nil, err
	}

	results := make([]scopePermissionResult.ScopePermissionRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = scopePermissionResult.ScopePermissionMapper(entity)
	}
	return results, nil
}

// GetScopePermissionsByPermissionID 根据权限ID获取作用域权限列表
func (s *Service) GetScopePermissionsByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]scopePermissionResult.ScopePermissionRO, error) {
	domainEntities, err := s.scopePermissionRepo.Get.ByPermissionID(ctx, permissionID)
	if err != nil {
		return nil, err
	}

	results := make([]scopePermissionResult.ScopePermissionRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = scopePermissionResult.ScopePermissionMapper(entity)
	}
	return results, nil
}
