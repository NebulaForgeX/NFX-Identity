package tenant_roles

import (
	"context"

	domain "nfxid/modules/access/domain/tenant_roles"
	"github.com/google/uuid"
)

// GetByID 按 ID 获取
func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*domain.TenantRole, error) {
	return s.repo.Get.ByID(ctx, id)
}

// GetByTenantIDAndRoleKey 按租户ID与角色键获取
func (s *Service) GetByTenantIDAndRoleKey(ctx context.Context, tenantID uuid.UUID, roleKey string) (*domain.TenantRole, error) {
	return s.repo.Get.ByTenantIDAndRoleKey(ctx, tenantID, roleKey)
}

// ListByTenantID 按租户ID列表
func (s *Service) ListByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*domain.TenantRole, error) {
	return s.repo.Get.ByTenantID(ctx, tenantID)
}
