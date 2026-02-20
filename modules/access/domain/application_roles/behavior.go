package application_roles

func (r *ApplicationRole) UpdateName(name *string) {
	r.state.Name = name
}

func (r *ApplicationRole) Update(roleKey string, name *string) error {
	if roleKey == "" {
		return ErrApplicationRoleRoleKeyRequired
	}
	r.state.RoleKey = roleKey
	r.state.Name = name
	return nil
}
