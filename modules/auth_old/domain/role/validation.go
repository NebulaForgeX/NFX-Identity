package role

import roleErrors "nfxid/modules/auth/domain/role/errors"

func (e *RoleEditable) Validate() error {
	if e.Name == "" {
		return roleErrors.ErrRoleNameRequired
	}
	if len(e.Name) > 50 {
		return roleErrors.ErrRoleNameRequired
	}
	return nil
}
