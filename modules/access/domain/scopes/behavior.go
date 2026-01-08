package scopes

import (
	"time"
)

func (s *Scope) Update(description *string) error {
	if s.DeletedAt() != nil {
		return ErrScopeNotFound
	}

	s.state.Description = description
	s.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (s *Scope) Delete() error {
	if s.IsSystem() {
		return ErrSystemScopeDelete
	}
	if s.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	s.state.DeletedAt = &now
	s.state.UpdatedAt = now
	return nil
}
