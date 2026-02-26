package ip_allowlist

import (
	clientsErr "nfxidentity/errors/src/clients"
	"time"

	"github.com/google/uuid"
)

type NewIPAllowlistParams struct {
	RuleID      string
	AppID       uuid.UUID
	CIDR        string
	Description *string
	Status      AllowlistStatus
	CreatedBy   *uuid.UUID
}

func NewIPAllowlist(p NewIPAllowlistParams) (*IPAllowlist, error) {
	if err := validateIPAllowlistParams(p); err != nil {
		return nil, err
	}

	status := p.Status
	if status == "" {
		status = AllowlistStatusActive
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewIPAllowlistFromState(IPAllowlistState{
		ID:          id,
		RuleID:      p.RuleID,
		AppID:       p.AppID,
		CIDR:        p.CIDR,
		Description: p.Description,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   p.CreatedBy,
	}), nil
}

func NewIPAllowlistFromState(st IPAllowlistState) *IPAllowlist {
	return &IPAllowlist{state: st}
}

func validateIPAllowlistParams(p NewIPAllowlistParams) error {
	if p.RuleID == "" {
		return clientsErr.ErrRuleIDRequired
	}
	if p.AppID == uuid.Nil {
		return clientsErr.ErrAppIDRequired
	}
	if p.CIDR == "" {
		return clientsErr.ErrCIDRRequired
	}
	if p.Status != "" {
		validStatuses := map[AllowlistStatus]struct{}{
			AllowlistStatusActive:   {},
			AllowlistStatusDisabled: {},
			AllowlistStatusRevoked:  {},
		}
		if _, ok := validStatuses[p.Status]; !ok {
			return clientsErr.ErrInvalidAllowlistStatus
		}
	}
	return nil
}
