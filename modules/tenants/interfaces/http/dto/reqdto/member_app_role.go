package reqdto

import (
	memberAppRoleAppCommands "nfxid/modules/tenants/application/member_app_roles/commands"

	"github.com/google/uuid"
)

type MemberAppRoleCreateRequestDTO struct {
	MemberID   uuid.UUID `json:"member_id" validate:"required,uuid"`
	AppID      uuid.UUID `json:"app_id" validate:"required,uuid"`
	RoleID     uuid.UUID `json:"role_id" validate:"required,uuid"`
	AssignedBy *uuid.UUID `json:"assigned_by,omitempty"`
	ExpiresAt  *string    `json:"expires_at,omitempty"`
}

type MemberAppRoleRevokeRequestDTO struct {
	ID           uuid.UUID `params:"id" validate:"required,uuid"`
	RevokedBy    uuid.UUID `json:"revoked_by" validate:"required,uuid"`
	RevokeReason string    `json:"revoke_reason"`
}

func (r *MemberAppRoleCreateRequestDTO) ToCreateCmd() memberAppRoleAppCommands.CreateMemberAppRoleCmd {
	return memberAppRoleAppCommands.CreateMemberAppRoleCmd{
		MemberID:   r.MemberID,
		AppID:      r.AppID,
		RoleID:     r.RoleID,
		AssignedBy: r.AssignedBy,
		ExpiresAt:  r.ExpiresAt,
	}
}

func (r *MemberAppRoleRevokeRequestDTO) ToRevokeCmd() memberAppRoleAppCommands.RevokeMemberAppRoleCmd {
	return memberAppRoleAppCommands.RevokeMemberAppRoleCmd{
		MemberAppRoleID: r.ID,
		RevokedBy:       r.RevokedBy,
		RevokeReason:    r.RevokeReason,
	}
}
