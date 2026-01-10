package mapper

import (
	memberGroupAppResult "nfxid/modules/tenants/application/member_groups/results"
	membergrouppb "nfxid/protos/gen/tenants/member_group"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// MemberGroupROToProto 将 MemberGroupRO 转换为 proto MemberGroup 消息
func MemberGroupROToProto(v *memberGroupAppResult.MemberGroupRO) *membergrouppb.MemberGroup {
	if v == nil {
		return nil
	}

	memberGroup := &membergrouppb.MemberGroup{
		Id:         v.ID.String(),
		MemberId:   v.MemberID.String(),
		GroupId:    v.GroupID.String(),
		AssignedAt: timestamppb.New(v.AssignedAt),
	}

	if v.AssignedBy != nil {
		assignedBy := v.AssignedBy.String()
		memberGroup.AssignedBy = &assignedBy
	}

	if v.RevokedAt != nil {
		memberGroup.RevokedAt = timestamppb.New(*v.RevokedAt)
	}

	if v.RevokedBy != nil {
		revokedBy := v.RevokedBy.String()
		memberGroup.RevokedBy = &revokedBy
	}

	return memberGroup
}

// MemberGroupListROToProto 批量转换 MemberGroupRO 到 proto MemberGroup
func MemberGroupListROToProto(results []memberGroupAppResult.MemberGroupRO) []*membergrouppb.MemberGroup {
	memberGroups := make([]*membergrouppb.MemberGroup, len(results))
	for i, v := range results {
		memberGroups[i] = MemberGroupROToProto(&v)
	}
	return memberGroups
}
