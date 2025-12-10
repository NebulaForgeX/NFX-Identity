package commands

import (
	"nebulaid/modules/auth/domain/badge"

	"github.com/google/uuid"
)

type CreateBadgeCmd struct {
	Editable badge.BadgeEditable
}

type UpdateBadgeCmd struct {
	BadgeID  uuid.UUID
	Editable badge.BadgeEditable
}
