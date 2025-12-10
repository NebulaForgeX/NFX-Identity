package views

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type UserBadgesView struct {
	UserID      uuid.UUID
	Username    string
	Email       string
	ProfileID   *uuid.UUID
	DisplayName *string
	Badges      *datatypes.JSON
}
