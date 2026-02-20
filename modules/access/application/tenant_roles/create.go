package tenant_roles

import (
	"context"

	domain "nfxid/modules/access/domain/tenant_roles"
	"github.com/google/uuid"
)

// CreateParams 创建参数
type CreateParams struct {
	TenantID uuid.UUID
	RoleKey  string
	Name     *string
}

// Create 创建租户角色
func (s *Service) Create(ctx context.Context, p CreateParams) (*domain.TenantRole, error) {
	exists, _ := s.repo.Check.ByTenantIDAndRoleKey(ctx, p.TenantID, p.RoleKey)
	if exists {
		return nil, domain.ErrTenantRoleKeyExistsInTenant
	}
	r, err := domain.NewTenantRole(domain.NewTenantRoleParams{
		TenantID: p.TenantID,
		RoleKey:  p.RoleKey,
		Name:     p.Name,
	})
	if err != nil {
		return nil, err
	}
	if err := s.repo.Create.New(ctx, r); err != nil {
		return nil, err
	}
	return r, nil
}
