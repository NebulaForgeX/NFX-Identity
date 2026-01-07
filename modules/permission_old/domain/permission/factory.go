package permission

import (
	"time"

	"github.com/google/uuid"
)

type NewPermissionParams struct {
	Editable PermissionEditable
	IsSystem bool
}

func NewPermission(p NewPermissionParams) (*Permission, error) {
	if err := p.Editable.Validate(); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewPermissionFromState(PermissionState{
		ID:        id,
		Editable:  p.Editable,
		IsSystem:  p.IsSystem,
		CreatedAt: now,
		UpdatedAt: now,
	}), nil
}

func NewPermissionFromState(st PermissionState) *Permission {
	return &Permission{state: st}
}

