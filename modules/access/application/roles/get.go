package roles

import (
	"context"
	roleResult "nfxid/modules/access/application/roles/results"

	"github.com/google/uuid"
)

// GetRole 根据ID获取角色
func (s *Service) GetRole(ctx context.Context, roleID uuid.UUID) (roleResult.RoleRO, error) {
	domainEntity, err := s.roleRepo.Get.ByID(ctx, roleID)
	if err != nil {
		return roleResult.RoleRO{}, err
	}
	return roleResult.RoleMapper(domainEntity), nil
}

// GetRoleByKey 根据Key获取角色
func (s *Service) GetRoleByKey(ctx context.Context, key string) (roleResult.RoleRO, error) {
	domainEntity, err := s.roleRepo.Get.ByKey(ctx, key)
	if err != nil {
		return roleResult.RoleRO{}, err
	}
	return roleResult.RoleMapper(domainEntity), nil
}
