package views

import (
	"time"

	"github.com/google/uuid"
)

type UserRoleView struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	RoleID    uuid.UUID
	CreatedAt time.Time
}

