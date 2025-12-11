package permission

import (
	"github.com/google/uuid"
)

type CreatePermissionCmd struct {
	Tag         string
	Name        string
	Description string
	Category    string
	IsSystem    bool
}

type UpdatePermissionCmd struct {
	ID          uuid.UUID
	Tag         string
	Name        string
	Description string
	Category    string
}

type DeletePermissionCmd struct {
	ID uuid.UUID
}

type GetPermissionCmd struct {
	ID uuid.UUID
}

type GetPermissionByTagCmd struct {
	Tag string
}

type ListPermissionsCmd struct {
	Category string
}

