package invitations

import "nfxid/pkgs/errx"

var (
	ErrInvitationNotFound       = errx.NotFound("INVITATION_NOT_FOUND", "invitation not found")
	ErrInviteIDRequired         = errx.InvalidArg("INVITE_ID_REQUIRED", "invite id is required")
	ErrTenantIDRequired         = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrEmailRequired            = errx.InvalidArg("EMAIL_REQUIRED", "email is required")
	ErrTokenHashRequired        = errx.InvalidArg("TOKEN_HASH_REQUIRED", "token hash is required")
	ErrExpiresAtRequired        = errx.InvalidArg("EXPIRES_AT_REQUIRED", "expires at is required")
	ErrInvitedByRequired        = errx.InvalidArg("INVITED_BY_REQUIRED", "invited by is required")
	ErrInviteIDAlreadyExists   = errx.Conflict("INVITE_ID_ALREADY_EXISTS", "invite id already exists")
	ErrInvalidInvitationStatus = errx.InvalidArg("INVALID_INVITATION_STATUS", "invalid invitation status")
	ErrInvitationExpired       = errx.Expired("INVITATION_EXPIRED", "invitation expired")
	ErrInvitationAlreadyAccepted = errx.Conflict("INVITATION_ALREADY_ACCEPTED", "invitation already accepted")
	ErrInvitationAlreadyRevoked = errx.Conflict("INVITATION_ALREADY_REVOKED", "invitation already revoked")
)
