package views

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type RoleView struct {
	ID          uuid.UUID
	Name        string
	Description *string
	Permissions *datatypes.JSON
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

