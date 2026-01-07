package user_role

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	state UserRoleState
}

type UserRoleState struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	RoleID    uuid.UUID
	CreatedAt time.Time
}

func (ur *UserRole) ID() uuid.UUID     { return ur.state.ID }
func (ur *UserRole) UserID() uuid.UUID { return ur.state.UserID }
func (ur *UserRole) RoleID() uuid.UUID { return ur.state.RoleID }
func (ur *UserRole) CreatedAt() time.Time { return ur.state.CreatedAt }

