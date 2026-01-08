package ip_allowlist

import "errors"

var (
	ErrIPAllowlistNotFound   = errors.New("ip allowlist not found")
	ErrRuleIDRequired        = errors.New("rule id is required")
	ErrAppIDRequired         = errors.New("app id is required")
	ErrCIDRRequired          = errors.New("cidr is required")
	ErrRuleIDAlreadyExists   = errors.New("rule id already exists")
	ErrInvalidAllowlistStatus = errors.New("invalid allowlist status")
	ErrIPAllowlistAlreadyRevoked = errors.New("ip allowlist already revoked")
)
