package user_role

import (
	"time"

	userRoleErrors "nfxid/modules/auth/domain/user_role/errors"

	"github.com/google/uuid"
)

type NewUserRoleParams struct {
	UserID uuid.UUID
	RoleID uuid.UUID
}

func NewUserRole(p NewUserRoleParams) (*UserRole, error) {
	if p.UserID == uuid.Nil {
		return nil, userRoleErrors.ErrUserIDRequired
	}
	if p.RoleID == uuid.Nil {
		return nil, userRoleErrors.ErrRoleIDRequired
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewUserRoleFromState(UserRoleState{
		ID:        id,
		UserID:    p.UserID,
		RoleID:    p.RoleID,
		CreatedAt: now,
	}), nil
}

func NewUserRoleFromState(st UserRoleState) *UserRole {
	return &UserRole{state: st}
}

