package permissions

import (
	"time"
)

func (p *Permission) Update(key, name string, description *string) error {
	if p.DeletedAt() != nil {
		return ErrPermissionNotFound
	}
	if key == "" {
		return ErrPermissionKeyRequired
	}
	if name == "" {
		return ErrPermissionNameRequired
	}

	p.state.Key = key
	p.state.Name = name
	p.state.Description = description
	p.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (p *Permission) Delete() error {
	if p.IsSystem() {
		return ErrSystemPermissionDelete
	}
	if p.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	p.state.DeletedAt = &now
	p.state.UpdatedAt = now
	return nil
}
