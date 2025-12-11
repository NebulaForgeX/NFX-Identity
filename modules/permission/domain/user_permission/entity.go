package user_permission

import (
	"time"

	"github.com/google/uuid"
)

type UserPermission struct {
	state UserPermissionState
}

type UserPermissionState struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	PermissionID uuid.UUID
	CreatedAt    time.Time
	DeletedAt    *time.Time
}

func (up *UserPermission) ID() uuid.UUID          { return up.state.ID }
func (up *UserPermission) UserID() uuid.UUID      { return up.state.UserID }
func (up *UserPermission) PermissionID() uuid.UUID { return up.state.PermissionID }
func (up *UserPermission) CreatedAt() time.Time  { return up.state.CreatedAt }
func (up *UserPermission) DeletedAt() *time.Time  { return up.state.DeletedAt }

func (up *UserPermission) IsActive() bool {
	return up.DeletedAt() == nil
}

