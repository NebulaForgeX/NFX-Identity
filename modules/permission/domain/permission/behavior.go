package permission

import (
	permissionErrors "nfxid/modules/permission/domain/permission/errors"
	"time"
)

func (p *Permission) EnsureEditable(e PermissionEditable) error {
	if err := e.Validate(); err != nil {
		return err
	}
	if p.DeletedAt() != nil {
		return permissionErrors.ErrPermissionNotFound
	}
	if p.IsSystem() {
		return permissionErrors.ErrPermissionSystemCannotModify
	}
	return nil
}

func (p *Permission) Update(e PermissionEditable) error {
	if err := p.EnsureEditable(e); err != nil {
		return err
	}

	p.state.Editable = e
	p.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (p *Permission) Delete() error {
	if p.DeletedAt() != nil {
		return nil // idempotent
	}
	if p.IsSystem() {
		return permissionErrors.ErrPermissionSystemCannotDelete
	}

	now := time.Now().UTC()
	p.state.DeletedAt = &now
	p.state.UpdatedAt = now
	return nil
}

