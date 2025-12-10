package views

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type RoleView struct {
	ID          uuid.UUID
	Name        string
	Description *string
	Permissions *datatypes.JSON
	IsSystem    bool
}
