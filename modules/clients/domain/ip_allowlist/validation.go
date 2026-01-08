package ip_allowlist

import "github.com/google/uuid"

func (ip *IPAllowlist) Validate() error {
	if ip.RuleID() == "" {
		return ErrRuleIDRequired
	}
	if ip.AppID() == uuid.Nil {
		return ErrAppIDRequired
	}
	if ip.CIDR() == "" {
		return ErrCIDRRequired
	}
	validStatuses := map[AllowlistStatus]struct{}{
		AllowlistStatusActive:   {},
		AllowlistStatusDisabled: {},
		AllowlistStatusRevoked:  {},
	}
	if _, ok := validStatuses[ip.Status()]; !ok {
		return ErrInvalidAllowlistStatus
	}
	return nil
}
