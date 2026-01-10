package mapper

import (
	memberAppRoleAppResult "nfxid/modules/tenants/application/member_app_roles/results"
	memberapprolepb "nfxid/protos/gen/tenants/member_app_role"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// MemberAppRoleROToProto 将 MemberAppRoleRO 转换为 proto MemberAppRole 消息
func MemberAppRoleROToProto(v *memberAppRoleAppResult.MemberAppRoleRO) *memberapprolepb.MemberAppRole {
	if v == nil {
		return nil
	}

	memberAppRole := &memberapprolepb.MemberAppRole{
		Id:         v.ID.String(),
		MemberId:   v.MemberID.String(),
		AppId:      v.AppID.String(),
		RoleId:     v.RoleID.String(),
		AssignedAt: timestamppb.New(v.AssignedAt),
	}

	if v.AssignedBy != nil {
		assignedBy := v.AssignedBy.String()
		memberAppRole.AssignedBy = &assignedBy
	}

	if v.ExpiresAt != nil {
		memberAppRole.ExpiresAt = timestamppb.New(*v.ExpiresAt)
	}

	if v.RevokedAt != nil {
		memberAppRole.RevokedAt = timestamppb.New(*v.RevokedAt)
	}

	if v.RevokedBy != nil {
		revokedBy := v.RevokedBy.String()
		memberAppRole.RevokedBy = &revokedBy
	}

	if v.RevokeReason != nil {
		memberAppRole.RevokeReason = v.RevokeReason
	}

	return memberAppRole
}

// MemberAppRoleListROToProto 批量转换 MemberAppRoleRO 到 proto MemberAppRole
func MemberAppRoleListROToProto(results []memberAppRoleAppResult.MemberAppRoleRO) []*memberapprolepb.MemberAppRole {
	memberAppRoles := make([]*memberapprolepb.MemberAppRole, len(results))
	for i, v := range results {
		memberAppRoles[i] = MemberAppRoleROToProto(&v)
	}
	return memberAppRoles
}
