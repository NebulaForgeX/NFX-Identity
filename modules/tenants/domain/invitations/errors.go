package invitations

import "errors"

var (
	ErrInvitationNotFound      = errors.New("invitation not found")
	ErrInviteIDRequired        = errors.New("invite id is required")
	ErrTenantIDRequired        = errors.New("tenant id is required")
	ErrEmailRequired           = errors.New("email is required")
	ErrTokenHashRequired       = errors.New("token hash is required")
	ErrExpiresAtRequired       = errors.New("expires at is required")
	ErrInvitedByRequired       = errors.New("invited by is required")
	ErrInviteIDAlreadyExists   = errors.New("invite id already exists")
	ErrInvalidInvitationStatus = errors.New("invalid invitation status")
	ErrInvitationExpired       = errors.New("invitation expired")
	ErrInvitationAlreadyAccepted = errors.New("invitation already accepted")
	ErrInvitationAlreadyRevoked = errors.New("invitation already revoked")
)
