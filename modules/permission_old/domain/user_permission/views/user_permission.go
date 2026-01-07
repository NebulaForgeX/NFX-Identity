package views

import (
	"nfxid/enums"
	"time"

	"github.com/google/uuid"
)

type UserPermissionView struct {
	ID           uuid.UUID                `json:"id"`
	UserID       uuid.UUID                `json:"user_id"`
	PermissionID uuid.UUID                `json:"permission_id"`
	Tag          string                   `json:"tag"`      // Permission tag for convenience
	Name         string                   `json:"name"`     // Permission name for convenience
	Category     enums.PermissionCategory `json:"category"` // Permission category for convenience
	CreatedAt    time.Time                `json:"created_at"`
}
