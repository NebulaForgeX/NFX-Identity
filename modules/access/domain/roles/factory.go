package roles

import (
	"time"

	"github.com/google/uuid"
)

type NewRoleParams struct {
	Key         string
	Name        string
	Description *string
	ScopeType   ScopeType
	IsSystem    bool
}

func NewRole(p NewRoleParams) (*Role, error) {
	if err := validateRoleParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	scopeType := p.ScopeType
	if scopeType == "" {
		scopeType = ScopeTypeTenant
	}

	now := time.Now().UTC()
	return NewRoleFromState(RoleState{
		ID:          id,
		Key:         p.Key,
		Name:        p.Name,
		Description: p.Description,
		ScopeType:   scopeType,
		IsSystem:    p.IsSystem,
		CreatedAt:   now,
		UpdatedAt:   now,
	}), nil
}

func NewRoleFromState(st RoleState) *Role {
	return &Role{state: st}
}

func validateRoleParams(p NewRoleParams) error {
	if p.Key == "" {
		return ErrRoleKeyRequired
	}
	if p.Name == "" {
		return ErrRoleNameRequired
	}
	if p.ScopeType != "" {
		validScopeTypes := map[ScopeType]struct{}{
			ScopeTypeTenant: {},
			ScopeTypeApp:    {},
			ScopeTypeGlobal: {},
		}
		if _, ok := validScopeTypes[p.ScopeType]; !ok {
			return ErrInvalidScopeType
		}
	}
	return nil
}
