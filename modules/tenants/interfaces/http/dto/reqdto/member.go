package reqdto

import (
	memberAppCommands "nfxid/modules/tenants/application/members/commands"
	memberDomain "nfxid/modules/tenants/domain/members"

	"github.com/google/uuid"
)

type MemberCreateRequestDTO struct {
	TenantID   uuid.UUID              `json:"tenant_id" validate:"required,uuid"`
	UserID     uuid.UUID              `json:"user_id" validate:"required,uuid"`
	Status     string                 `json:"status,omitempty"`
	Source     string                 `json:"source,omitempty"`
	CreatedBy  *uuid.UUID             `json:"created_by,omitempty"`
	ExternalRef *string                `json:"external_ref,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

type MemberUpdateStatusRequestDTO struct {
	ID     uuid.UUID `params:"id" validate:"required,uuid"`
	Status string    `json:"status" validate:"required"`
}

type MemberByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

func (r *MemberCreateRequestDTO) ToCreateCmd() memberAppCommands.CreateMemberCmd {
	cmd := memberAppCommands.CreateMemberCmd{
		TenantID:    r.TenantID,
		UserID:      r.UserID,
		CreatedBy:   r.CreatedBy,
		ExternalRef: r.ExternalRef,
		Metadata:    r.Metadata,
	}

	// Parse status
	if r.Status != "" {
		cmd.Status = memberDomain.MemberStatus(r.Status)
	} else {
		cmd.Status = memberDomain.MemberStatusInvited
	}

	// Parse source
	if r.Source != "" {
		cmd.Source = memberDomain.MemberSource(r.Source)
	} else {
		cmd.Source = memberDomain.MemberSourceManual
	}

	return cmd
}

func (r *MemberUpdateStatusRequestDTO) ToUpdateStatusCmd() memberAppCommands.UpdateMemberStatusCmd {
	return memberAppCommands.UpdateMemberStatusCmd{
		MemberID: r.ID,
		Status:   memberDomain.MemberStatus(r.Status),
	}
}
