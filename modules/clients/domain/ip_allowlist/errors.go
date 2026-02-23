package ip_allowlist

import "nfxid/pkgs/errx"

var (
	ErrIPAllowlistNotFound       = errx.NotFound("IP_ALLOWLIST_NOT_FOUND", "ip allowlist not found")
	ErrRuleIDRequired            = errx.InvalidArg("RULE_ID_REQUIRED", "rule id is required")
	ErrAppIDRequired             = errx.InvalidArg("APP_ID_REQUIRED", "app id is required")
	ErrCIDRRequired              = errx.InvalidArg("CIDR_REQUIRED", "cidr is required")
	ErrRuleIDAlreadyExists       = errx.Conflict("RULE_ID_ALREADY_EXISTS", "rule id already exists")
	ErrInvalidAllowlistStatus    = errx.InvalidArg("INVALID_ALLOWLIST_STATUS", "invalid allowlist status")
	ErrIPAllowlistAlreadyRevoked = errx.Conflict("IP_ALLOWLIST_ALREADY_REVOKED", "ip allowlist already revoked")
)
