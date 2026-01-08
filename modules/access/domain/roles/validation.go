package roles

func (r *Role) Validate() error {
	if r.Key() == "" {
		return ErrRoleKeyRequired
	}
	if r.Name() == "" {
		return ErrRoleNameRequired
	}
	validScopeTypes := map[ScopeType]struct{}{
		ScopeTypeTenant: {},
		ScopeTypeApp:    {},
		ScopeTypeGlobal: {},
	}
	if _, ok := validScopeTypes[r.ScopeType()]; !ok {
		return ErrInvalidScopeType
	}
	return nil
}
