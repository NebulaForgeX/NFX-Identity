package ip_allowlist

import (
	clientsErr "nfxidentity/errors/src/clients"

	"github.com/google/uuid"
)

func (ip *IPAllowlist) Validate() error {
	if ip.RuleID() == "" {
		return clientsErr.ErrRuleIDRequired
	}
	if ip.AppID() == uuid.Nil {
		return clientsErr.ErrAppIDRequired
	}
	if ip.CIDR() == "" {
		return clientsErr.ErrCIDRRequired
	}
	validStatuses := map[AllowlistStatus]struct{}{
		AllowlistStatusActive:   {},
		AllowlistStatusDisabled: {},
		AllowlistStatusRevoked:  {},
	}
	if _, ok := validStatuses[ip.Status()]; !ok {
		return clientsErr.ErrInvalidAllowlistStatus
	}
	return nil
}
