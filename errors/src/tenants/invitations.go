package tenants

import "nfxid/pkgs/errx"

const (
	CodeInvitationNotFound        = "INVITATION_NOT_FOUND"
	CodeInviteIDRequired          = "INVITE_ID_REQUIRED"
	CodeEmailRequired             = "EMAIL_REQUIRED"
	CodeTokenHashRequired         = "TOKEN_HASH_REQUIRED"
	CodeInvitedByRequired         = "INVITED_BY_REQUIRED"
	CodeInviteIDAlreadyExists     = "INVITE_ID_ALREADY_EXISTS"
	CodeInvalidInvitationStatus   = "INVALID_INVITATION_STATUS"
	CodeInvitationExpired         = "INVITATION_EXPIRED"
	CodeInvitationAlreadyAccepted = "INVITATION_ALREADY_ACCEPTED"
	CodeInvitationAlreadyRevoked  = "INVITATION_ALREADY_REVOKED"
)

var (
	ErrInvitationNotFound        = errx.NotFound(CodeInvitationNotFound, "invitation not found")
	ErrInviteIDRequired          = errx.InvalidArg(CodeInviteIDRequired, "invite id is required")
	ErrEmailRequired             = errx.InvalidArg(CodeEmailRequired, "email is required")
	ErrTokenHashRequired         = errx.InvalidArg(CodeTokenHashRequired, "token hash is required")
	ErrInvitedByRequired         = errx.InvalidArg(CodeInvitedByRequired, "invited by is required")
	ErrInviteIDAlreadyExists     = errx.Conflict(CodeInviteIDAlreadyExists, "invite id already exists")
	ErrInvalidInvitationStatus   = errx.InvalidArg(CodeInvalidInvitationStatus, "invalid invitation status")
	ErrInvitationExpired         = errx.Expired(CodeInvitationExpired, "invitation expired")
	ErrInvitationAlreadyAccepted = errx.Conflict(CodeInvitationAlreadyAccepted, "invitation already accepted")
	ErrInvitationAlreadyRevoked  = errx.Conflict(CodeInvitationAlreadyRevoked, "invitation already revoked")
)
