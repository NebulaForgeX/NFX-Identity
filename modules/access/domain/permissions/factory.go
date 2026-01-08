package permissions

import (
	"time"

	"github.com/google/uuid"
)

type NewPermissionParams struct {
	Key         string
	Name        string
	Description *string
	IsSystem    bool
}

func NewPermission(p NewPermissionParams) (*Permission, error) {
	if err := validatePermissionParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewPermissionFromState(PermissionState{
		ID:          id,
		Key:         p.Key,
		Name:        p.Name,
		Description: p.Description,
		IsSystem:    p.IsSystem,
		CreatedAt:   now,
		UpdatedAt:   now,
	}), nil
}

func NewPermissionFromState(st PermissionState) *Permission {
	return &Permission{state: st}
}

func validatePermissionParams(p NewPermissionParams) error {
	if p.Key == "" {
		return ErrPermissionKeyRequired
	}
	if p.Name == "" {
		return ErrPermissionNameRequired
	}
	return nil
}
