package role_permissions

import (
	"time"

	"github.com/google/uuid"
)

type NewRolePermissionParams struct {
	RoleID       uuid.UUID
	PermissionID uuid.UUID
	CreatedBy    *uuid.UUID
}

func NewRolePermission(p NewRolePermissionParams) (*RolePermission, error) {
	if err := validateRolePermissionParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewRolePermissionFromState(RolePermissionState{
		ID:           id,
		RoleID:       p.RoleID,
		PermissionID: p.PermissionID,
		CreatedAt:    now,
		CreatedBy:    p.CreatedBy,
	}), nil
}

func NewRolePermissionFromState(st RolePermissionState) *RolePermission {
	return &RolePermission{state: st}
}

func validateRolePermissionParams(p NewRolePermissionParams) error {
	if p.RoleID == uuid.Nil {
		return ErrRoleIDRequired
	}
	if p.PermissionID == uuid.Nil {
		return ErrPermissionIDRequired
	}
	return nil
}
