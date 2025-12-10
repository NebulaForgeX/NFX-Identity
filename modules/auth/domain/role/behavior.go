package role

import (
	"time"

	roleErrors "nebulaid/modules/auth/domain/role/errors"
)

func (r *Role) EnsureEditable(e RoleEditable) error {
	if err := e.Validate(); err != nil {
		return err
	}
	if r.DeletedAt() != nil {
		return roleErrors.ErrRoleNotFound
	}
	return nil
}

func (r *Role) Update(e RoleEditable) error {
	if err := r.EnsureEditable(e); err != nil {
		return err
	}

	r.state.Editable = e
	r.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (r *Role) Delete() error {
	if r.IsSystem() {
		return roleErrors.ErrSystemRoleDelete
	}
	if r.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	r.state.DeletedAt = &now
	r.state.UpdatedAt = now
	return nil
}
