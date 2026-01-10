package results

import (
	"time"

	"nfxid/modules/clients/domain/ip_allowlist"

	"github.com/google/uuid"
)

type IPAllowlistRO struct {
	ID          uuid.UUID
	RuleID      string
	AppID       uuid.UUID
	CIDR        string
	Description *string
	Status      ip_allowlist.AllowlistStatus
	CreatedAt   time.Time
	CreatedBy   *uuid.UUID
	UpdatedAt   time.Time
	UpdatedBy   *uuid.UUID
	RevokedAt   *time.Time
	RevokedBy   *uuid.UUID
	RevokeReason *string
}

// IPAllowlistMapper 将 Domain IPAllowlist 转换为 Application IPAllowlistRO
func IPAllowlistMapper(ip *ip_allowlist.IPAllowlist) IPAllowlistRO {
	if ip == nil {
		return IPAllowlistRO{}
	}

	return IPAllowlistRO{
		ID:          ip.ID(),
		RuleID:      ip.RuleID(),
		AppID:       ip.AppID(),
		CIDR:        ip.CIDR(),
		Description: ip.Description(),
		Status:      ip.Status(),
		CreatedAt:   ip.CreatedAt(),
		CreatedBy:   ip.CreatedBy(),
		UpdatedAt:   ip.UpdatedAt(),
		UpdatedBy:   ip.UpdatedBy(),
		RevokedAt:   ip.RevokedAt(),
		RevokedBy:   ip.RevokedBy(),
		RevokeReason: ip.RevokeReason(),
	}
}
