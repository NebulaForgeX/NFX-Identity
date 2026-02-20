package mapper

import (
	"nfxid/modules/access/domain/tenant_roles"
	"nfxid/modules/access/infrastructure/rdb/models"
)

func TenantRoleDomainToModel(r *tenant_roles.TenantRole) *models.TenantRole {
	if r == nil {
		return nil
	}
	return &models.TenantRole{
		ID: r.ID(), TenantID: r.TenantID(), RoleKey: r.RoleKey(),
		Name: r.Name(), CreatedAt: r.CreatedAt(),
	}
}

func TenantRoleModelToDomain(m *models.TenantRole) *tenant_roles.TenantRole {
	if m == nil {
		return nil
	}
	return tenant_roles.NewTenantRoleFromState(tenant_roles.TenantRoleState{
		ID: m.ID, TenantID: m.TenantID, RoleKey: m.RoleKey,
		Name: m.Name, CreatedAt: m.CreatedAt,
	})
}
