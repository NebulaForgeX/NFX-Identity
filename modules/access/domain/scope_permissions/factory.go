package scope_permissions

import (
	"time"

	"github.com/google/uuid"
)

type NewScopePermissionParams struct {
	Scope        string
	PermissionID uuid.UUID
}

func NewScopePermission(p NewScopePermissionParams) (*ScopePermission, error) {
	if err := validateScopePermissionParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewScopePermissionFromState(ScopePermissionState{
		ID:           id,
		Scope:        p.Scope,
		PermissionID: p.PermissionID,
		CreatedAt:    now,
	}), nil
}

func NewScopePermissionFromState(st ScopePermissionState) *ScopePermission {
	return &ScopePermission{state: st}
}

func validateScopePermissionParams(p NewScopePermissionParams) error {
	if p.Scope == "" {
		return ErrScopeRequired
	}
	if p.PermissionID == uuid.Nil {
		return ErrPermissionIDRequired
	}
	return nil
}
