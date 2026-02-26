package reqdto

import (
	memberGroupAppCommands "nfxidentity/modules/tenants/application/member_groups/commands"

	"github.com/google/uuid"
)

type MemberGroupCreateRequestDTO struct {
	MemberID   uuid.UUID  `json:"member_id"             validate:"required,uuid"`
	GroupID    uuid.UUID  `json:"group_id"              validate:"required,uuid"`
	AssignedBy *uuid.UUID `json:"assigned_by,omitempty"`
}

type MemberGroupRevokeRequestDTO struct {
	ID        uuid.UUID `uri:"id" validate:"required,uuid"`
	RevokedBy uuid.UUID `         validate:"required,uuid" json:"revoked_by"`
}

func (r *MemberGroupCreateRequestDTO) ToCreateCmd() memberGroupAppCommands.CreateMemberGroupCmd {
	return memberGroupAppCommands.CreateMemberGroupCmd{
		MemberID:   r.MemberID,
		GroupID:    r.GroupID,
		AssignedBy: r.AssignedBy,
	}
}

func (r *MemberGroupRevokeRequestDTO) ToRevokeCmd() memberGroupAppCommands.RevokeMemberGroupCmd {
	return memberGroupAppCommands.RevokeMemberGroupCmd{
		MemberGroupID: r.ID,
		RevokedBy:     r.RevokedBy,
	}
}
