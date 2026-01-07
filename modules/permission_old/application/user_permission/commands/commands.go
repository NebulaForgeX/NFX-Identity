package user_permission

import "github.com/google/uuid"

type AssignPermissionCmd struct {
	UserID       uuid.UUID
	PermissionID uuid.UUID
}

type RevokePermissionCmd struct {
	UserID       uuid.UUID
	PermissionID uuid.UUID
}

type GetUserPermissionsCmd struct {
	UserID uuid.UUID
}

type CheckPermissionCmd struct {
	UserID uuid.UUID
	Tag    string
}

