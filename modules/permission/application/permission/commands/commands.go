package permission

import (
	"nfxid/enums"

	"github.com/google/uuid"
)

type CreatePermissionCmd struct {
	Tag         string
	Name        string
	Description string
	Category    enums.PermissionCategory
	IsSystem    bool
}

type UpdatePermissionCmd struct {
	ID          uuid.UUID
	Tag         string
	Name        string
	Description string
	Category    enums.PermissionCategory
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
	Category enums.PermissionCategory
}
