package user_permission

import (
	"time"

	userPermissionErrors "nfxid/modules/permission/domain/user_permission/errors"

	"github.com/google/uuid"
)

type NewUserPermissionParams struct {
	UserID       uuid.UUID
	PermissionID uuid.UUID
}

func NewUserPermission(p NewUserPermissionParams) (*UserPermission, error) {
	if p.UserID == uuid.Nil {
		return nil, userPermissionErrors.ErrUserPermissionUserIDRequired
	}
	if p.PermissionID == uuid.Nil {
		return nil, userPermissionErrors.ErrUserPermissionPermissionIDRequired
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewUserPermissionFromState(UserPermissionState{
		ID:           id,
		UserID:       p.UserID,
		PermissionID: p.PermissionID,
		CreatedAt:    now,
	}), nil
}

func NewUserPermissionFromState(st UserPermissionState) *UserPermission {
	return &UserPermission{state: st}
}

