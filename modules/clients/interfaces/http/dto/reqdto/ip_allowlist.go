package reqdto

import (
	ipAllowlistAppCommands "nfxid/modules/clients/application/ip_allowlist/commands"
	ipAllowlistDomain "nfxid/modules/clients/domain/ip_allowlist"

	"github.com/google/uuid"
)

type IPAllowlistCreateRequestDTO struct {
	RuleID      string     `json:"rule_id" validate:"required"`
	AppID       uuid.UUID  `json:"app_id" validate:"required"`
	CIDR        string     `json:"cidr" validate:"required"`
	Description *string    `json:"description,omitempty"`
	Status      string     `json:"status,omitempty"`
	CreatedBy   *uuid.UUID `json:"created_by,omitempty"`
}

type IPAllowlistByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type IPAllowlistDeleteRequestDTO struct {
	RuleID string `params:"rule_id" validate:"required"`
}

func (r *IPAllowlistCreateRequestDTO) ToCreateCmd() ipAllowlistAppCommands.CreateIPAllowlistCmd {
	cmd := ipAllowlistAppCommands.CreateIPAllowlistCmd{
		RuleID:      r.RuleID,
		AppID:       r.AppID,
		CIDR:        r.CIDR,
		Description: r.Description,
		CreatedBy:   r.CreatedBy,
	}

	if r.Status != "" {
		cmd.Status = ipAllowlistDomain.AllowlistStatus(r.Status)
	}

	return cmd
}
