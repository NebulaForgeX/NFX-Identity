package commands

import (
	"time"

	"github.com/google/uuid"
)

type CreateAuthorizationCodeCmd struct {
	Code      string
	MaxUses   int
	CreatedBy *uuid.UUID
	ExpiresAt *time.Time
	IsActive  bool
}

type UseAuthorizationCodeCmd struct {
	Code string
}

type GetAuthorizationCodeCmd struct {
	ID uuid.UUID
}

type GetAuthorizationCodeByCodeCmd struct {
	Code string
}

type DeleteAuthorizationCodeCmd struct {
	ID uuid.UUID
}

type ActivateAuthorizationCodeCmd struct {
	ID uuid.UUID
}

type DeactivateAuthorizationCodeCmd struct {
	ID uuid.UUID
}
