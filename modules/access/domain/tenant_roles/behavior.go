package tenant_roles

import (
	accessErr "nfxid/errors/src/access"
)

// UpdateName 更新名称
func (r *TenantRole) UpdateName(name *string) {
	r.state.Name = name
}

// Update 更新可写字段
func (r *TenantRole) Update(roleKey string, name *string) error {
	if roleKey == "" {
		return accessErr.ErrTenantRoleRoleKeyRequired
	}
	r.state.RoleKey = roleKey
	r.state.Name = name
	return nil
}
