package audit

import "nfxid/pkgs/errx"

const (
	CodeEventRetentionPolicyNotFound = "EVENT_RETENTION_POLICY_NOT_FOUND"
	CodePolicyNameRequired           = "POLICY_NAME_REQUIRED"
	CodeRetentionDaysRequired        = "RETENTION_DAYS_REQUIRED"
	CodeRetentionActionRequired      = "RETENTION_ACTION_REQUIRED"
	CodePolicyNameAlreadyExists      = "POLICY_NAME_ALREADY_EXISTS"
	CodeInvalidRetentionAction       = "INVALID_RETENTION_ACTION"
	CodeInvalidStatus                = "INVALID_STATUS"
)

var (
	ErrEventRetentionPolicyNotFound = errx.NotFound(CodeEventRetentionPolicyNotFound, "event retention policy not found")
	ErrPolicyNameRequired           = errx.InvalidArg(CodePolicyNameRequired, "policy name is required")
	ErrRetentionDaysRequired        = errx.InvalidArg(CodeRetentionDaysRequired, "retention days is required")
	ErrRetentionActionRequired      = errx.InvalidArg(CodeRetentionActionRequired, "retention action is required")
	ErrPolicyNameAlreadyExists      = errx.Conflict(CodePolicyNameAlreadyExists, "policy name already exists")
	ErrInvalidRetentionAction       = errx.InvalidArg(CodeInvalidRetentionAction, "invalid retention action")
	ErrInvalidStatus                = errx.InvalidArg(CodeInvalidStatus, "invalid status")
)
