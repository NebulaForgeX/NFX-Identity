package commands

import (
	"github.com/google/uuid"
)

type CreateUserRoleCmd struct {
	UserID uuid.UUID
	RoleID uuid.UUID
}

type DeleteUserRoleCmd struct {
	UserRoleID uuid.UUID
}

type DeleteUserRoleByUserAndRoleCmd struct {
	UserID uuid.UUID
	RoleID uuid.UUID
}

