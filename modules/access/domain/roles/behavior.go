package roles

import (
	"time"
)

func (r *Role) Update(key, name string, description *string, scopeType ScopeType) error {
	if r.DeletedAt() != nil {
		return ErrRoleNotFound
	}
	if key == "" {
		return ErrRoleKeyRequired
	}
	if name == "" {
		return ErrRoleNameRequired
	}
	if scopeType != "" {
		validScopeTypes := map[ScopeType]struct{}{
			ScopeTypeTenant: {},
			ScopeTypeApp:    {},
			ScopeTypeGlobal: {},
		}
		if _, ok := validScopeTypes[scopeType]; !ok {
			return ErrInvalidScopeType
		}
		r.state.ScopeType = scopeType
	}

	r.state.Key = key
	r.state.Name = name
	r.state.Description = description
	r.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (r *Role) Delete() error {
	if r.IsSystem() {
		return ErrSystemRoleDelete
	}
	if r.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	r.state.DeletedAt = &now
	r.state.UpdatedAt = now
	return nil
}
