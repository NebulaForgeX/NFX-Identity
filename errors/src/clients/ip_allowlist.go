package clients

import "nfxid/pkgs/errx"

const (
	CodeIPAllowlistNotFound       = "IP_ALLOWLIST_NOT_FOUND"
	CodeRuleIDRequired            = "RULE_ID_REQUIRED"
	CodeCIDRRequired              = "CIDR_REQUIRED"
	CodeRuleIDAlreadyExists       = "RULE_ID_ALREADY_EXISTS"
	CodeInvalidAllowlistStatus    = "INVALID_ALLOWLIST_STATUS"
	CodeIPAllowlistAlreadyRevoked = "IP_ALLOWLIST_ALREADY_REVOKED"
)

var (
	ErrIPAllowlistNotFound       = errx.NotFound(CodeIPAllowlistNotFound, "ip allowlist not found")
	ErrRuleIDRequired            = errx.InvalidArg(CodeRuleIDRequired, "rule id is required")
	ErrCIDRRequired              = errx.InvalidArg(CodeCIDRRequired, "cidr is required")
	ErrRuleIDAlreadyExists       = errx.Conflict(CodeRuleIDAlreadyExists, "rule id already exists")
	ErrInvalidAllowlistStatus    = errx.InvalidArg(CodeInvalidAllowlistStatus, "invalid allowlist status")
	ErrIPAllowlistAlreadyRevoked = errx.Conflict(CodeIPAllowlistAlreadyRevoked, "ip allowlist already revoked")
)
