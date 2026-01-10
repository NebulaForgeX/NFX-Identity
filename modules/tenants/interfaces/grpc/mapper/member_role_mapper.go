package mapper

import (
	memberRoleAppResult "nfxid/modules/tenants/application/member_roles/results"
	memberrolepb "nfxid/protos/gen/tenants/member_role"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// MemberRoleROToProto 将 MemberRoleRO 转换为 proto MemberRole 消息
func MemberRoleROToProto(v *memberRoleAppResult.MemberRoleRO) *memberrolepb.MemberRole {
	if v == nil {
		return nil
	}

	memberRole := &memberrolepb.MemberRole{
		Id:         v.ID.String(),
		TenantId:   v.TenantID.String(),
		MemberId:   v.MemberID.String(),
		RoleId:     v.RoleID.String(),
		AssignedAt: timestamppb.New(v.AssignedAt),
	}

	if v.AssignedBy != nil {
		assignedBy := v.AssignedBy.String()
		memberRole.AssignedBy = &assignedBy
	}

	if v.ExpiresAt != nil {
		memberRole.ExpiresAt = timestamppb.New(*v.ExpiresAt)
	}

	if v.Scope != nil {
		memberRole.Scope = v.Scope
	}

	if v.RevokedAt != nil {
		memberRole.RevokedAt = timestamppb.New(*v.RevokedAt)
	}

	if v.RevokedBy != nil {
		revokedBy := v.RevokedBy.String()
		memberRole.RevokedBy = &revokedBy
	}

	if v.RevokeReason != nil {
		memberRole.RevokeReason = v.RevokeReason
	}

	return memberRole
}

// MemberRoleListROToProto 批量转换 MemberRoleRO 到 proto MemberRole
func MemberRoleListROToProto(results []memberRoleAppResult.MemberRoleRO) []*memberrolepb.MemberRole {
	memberRoles := make([]*memberrolepb.MemberRole, len(results))
	for i, v := range results {
		memberRoles[i] = MemberRoleROToProto(&v)
	}
	return memberRoles
}
