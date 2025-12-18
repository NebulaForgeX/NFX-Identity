package views

import (
	"time"

	"github.com/google/uuid"
)

type UserView struct {
	ID           uuid.UUID
	Username     string
	Email        string
	Phone        *string
	PasswordHash string
	Status       string
	IsVerified   bool
	LastLoginAt  *time.Time
	Roles        []byte // JSON array of roles from user_with_role_view
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
