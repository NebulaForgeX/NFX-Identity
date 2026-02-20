package tenant_roles

// UpdateName 更新名称
func (r *TenantRole) UpdateName(name *string) {
	r.state.Name = name
}

// Update 更新可写字段
func (r *TenantRole) Update(roleKey string, name *string) error {
	if roleKey == "" {
		return ErrTenantRoleRoleKeyRequired
	}
	r.state.RoleKey = roleKey
	r.state.Name = name
	return nil
}
