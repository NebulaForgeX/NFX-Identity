package event_retention_policies

import "nfxid/pkgs/errx"

var (
	ErrEventRetentionPolicyNotFound = errx.NotFound("EVENT_RETENTION_POLICY_NOT_FOUND", "event retention policy not found")
	ErrPolicyNameRequired           = errx.InvalidArg("POLICY_NAME_REQUIRED", "policy name is required")
	ErrRetentionDaysRequired       = errx.InvalidArg("RETENTION_DAYS_REQUIRED", "retention days is required")
	ErrRetentionActionRequired     = errx.InvalidArg("RETENTION_ACTION_REQUIRED", "retention action is required")
	ErrPolicyNameAlreadyExists      = errx.Conflict("POLICY_NAME_ALREADY_EXISTS", "policy name already exists")
	ErrInvalidRetentionAction      = errx.InvalidArg("INVALID_RETENTION_ACTION", "invalid retention action")
	ErrInvalidStatus                = errx.InvalidArg("INVALID_STATUS", "invalid status")
)
