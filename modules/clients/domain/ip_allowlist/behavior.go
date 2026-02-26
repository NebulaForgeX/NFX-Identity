package ip_allowlist

import (
	clientsErr "nfxidentity/errors/src/clients"
	"time"

	"github.com/google/uuid"
)

func (ip *IPAllowlist) Update(description *string, cidr string, updatedBy *uuid.UUID) error {
	if ip.RevokedAt() != nil {
		return clientsErr.ErrIPAllowlistAlreadyRevoked
	}
	if cidr != "" {
		ip.state.CIDR = cidr
	}
	if description != nil {
		ip.state.Description = description
	}
	ip.state.UpdatedBy = updatedBy
	ip.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ip *IPAllowlist) UpdateStatus(status AllowlistStatus, updatedBy *uuid.UUID) error {
	if ip.RevokedAt() != nil {
		return clientsErr.ErrIPAllowlistAlreadyRevoked
	}
	validStatuses := map[AllowlistStatus]struct{}{
		AllowlistStatusActive:   {},
		AllowlistStatusDisabled: {},
		AllowlistStatusRevoked:  {},
	}
	if _, ok := validStatuses[status]; !ok {
		return clientsErr.ErrInvalidAllowlistStatus
	}

	ip.state.Status = status
	ip.state.UpdatedBy = updatedBy
	ip.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ip *IPAllowlist) Revoke(revokedBy uuid.UUID, reason string) error {
	if ip.RevokedAt() != nil {
		return clientsErr.ErrIPAllowlistAlreadyRevoked
	}

	now := time.Now().UTC()
	ip.state.RevokedAt = &now
	ip.state.RevokedBy = &revokedBy
	ip.state.RevokeReason = &reason
	ip.state.Status = AllowlistStatusRevoked
	ip.state.UpdatedAt = now
	return nil
}

func (ip *IPAllowlist) IsActive() bool {
	return ip.Status() == AllowlistStatusActive && ip.RevokedAt() == nil
}
