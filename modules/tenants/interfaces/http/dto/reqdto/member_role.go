package reqdto

import (
	memberRoleAppCommands "nfxid/modules/tenants/application/member_roles/commands"

	"github.com/google/uuid"
)

type MemberRoleCreateRequestDTO struct {
	TenantID   uuid.UUID  `json:"tenant_id" validate:"required,uuid"`
	MemberID   uuid.UUID  `json:"member_id" validate:"required,uuid"`
	RoleID     uuid.UUID  `json:"role_id" validate:"required,uuid"`
	AssignedBy *uuid.UUID `json:"assigned_by,omitempty"`
	ExpiresAt  *string    `json:"expires_at,omitempty"`
	Scope      *string    `json:"scope,omitempty"`
}

type MemberRoleRevokeRequestDTO struct {
	ID           uuid.UUID `uri:"id" validate:"required,uuid"`
	RevokedBy    uuid.UUID `json:"revoked_by" validate:"required,uuid"`
	RevokeReason string    `json:"revoke_reason"`
}

func (r *MemberRoleCreateRequestDTO) ToCreateCmd() memberRoleAppCommands.CreateMemberRoleCmd {
	return memberRoleAppCommands.CreateMemberRoleCmd{
		TenantID:   r.TenantID,
		MemberID:   r.MemberID,
		RoleID:     r.RoleID,
		AssignedBy: r.AssignedBy,
		ExpiresAt:  r.ExpiresAt,
		Scope:      r.Scope,
	}
}

func (r *MemberRoleRevokeRequestDTO) ToRevokeCmd() memberRoleAppCommands.RevokeMemberRoleCmd {
	return memberRoleAppCommands.RevokeMemberRoleCmd{
		MemberRoleID: r.ID,
		RevokedBy:    r.RevokedBy,
		RevokeReason: r.RevokeReason,
	}
}
