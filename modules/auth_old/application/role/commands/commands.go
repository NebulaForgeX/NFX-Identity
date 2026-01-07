package commands

import (
	roleDomain "nfxid/modules/auth/domain/role"

	"github.com/google/uuid"
)

type CreateRoleCmd struct {
	Editable roleDomain.RoleEditable
	IsSystem bool
}

type UpdateRoleCmd struct {
	RoleID   uuid.UUID
	Editable roleDomain.RoleEditable
}
