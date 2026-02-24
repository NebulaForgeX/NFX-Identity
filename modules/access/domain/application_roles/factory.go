package application_roles

import (
	accessErr "nfxid/errors/src/access"
	"time"

	"github.com/google/uuid"
)

type NewApplicationRoleParams struct {
	ApplicationID uuid.UUID
	RoleKey       string
	Name          *string
}

func NewApplicationRole(p NewApplicationRoleParams) (*ApplicationRole, error) {
	if err := validateNewParams(p); err != nil {
		return nil, err
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return NewApplicationRoleFromState(ApplicationRoleState{
		ID:            id,
		ApplicationID: p.ApplicationID,
		RoleKey:       p.RoleKey,
		Name:          p.Name,
		CreatedAt:     time.Now().UTC(),
	}), nil
}

func NewApplicationRoleFromState(st ApplicationRoleState) *ApplicationRole {
	return &ApplicationRole{state: st}
}

func validateNewParams(p NewApplicationRoleParams) error {
	if p.ApplicationID == uuid.Nil {
		return accessErr.ErrApplicationRoleApplicationIDRequired
	}
	if p.RoleKey == "" {
		return accessErr.ErrApplicationRoleRoleKeyRequired
	}
	return nil
}
