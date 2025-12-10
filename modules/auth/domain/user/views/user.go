package views

import (
	"time"

	"github.com/google/uuid"
)

type UserView struct {
	ID          uuid.UUID
	Username    string
	Email       string
	Phone       string
	Status      string
	IsVerified  bool
	LastLoginAt *time.Time
	RoleID      *uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
