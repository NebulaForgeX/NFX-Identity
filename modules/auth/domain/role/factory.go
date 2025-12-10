package role

import (
	"time"

	"github.com/google/uuid"
)

type NewRoleParams struct {
	Editable RoleEditable
	IsSystem bool
}

func NewRole(p NewRoleParams) (*Role, error) {
	if err := p.Editable.Validate(); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewRoleFromState(RoleState{
		ID:        id,
		Editable:  p.Editable,
		IsSystem:  p.IsSystem,
		CreatedAt: now,
		UpdatedAt: now,
	}), nil
}

func NewRoleFromState(st RoleState) *Role {
	return &Role{state: st}
}

