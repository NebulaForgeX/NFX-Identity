package mapper

import (
	memberAppResult "nfxid/modules/tenants/application/members/results"
	memberDomain "nfxid/modules/tenants/domain/members"
	memberpb "nfxid/protos/gen/tenants/member"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MemberROToProto 将 MemberRO 转换为 proto Member 消息
func MemberROToProto(v *memberAppResult.MemberRO) *memberpb.Member {
	if v == nil {
		return nil
	}

	member := &memberpb.Member{
		Id:        v.ID.String(),
		MemberId:  v.MemberID.String(),
		TenantId:  v.TenantID.String(),
		UserId:    v.UserID.String(),
		Status:    memberStatusToProto(v.Status),
		Source:    memberSourceToProto(v.Source),
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.JoinedAt != nil {
		member.JoinedAt = timestamppb.New(*v.JoinedAt)
	}

	if v.LeftAt != nil {
		member.LeftAt = timestamppb.New(*v.LeftAt)
	}

	if v.CreatedBy != nil {
		createdBy := v.CreatedBy.String()
		member.CreatedBy = &createdBy
	}

	if v.ExternalRef != nil {
		member.ExternalRef = v.ExternalRef
	}

	if v.Metadata != nil && len(v.Metadata) > 0 {
		if metadataStruct, err := structpb.NewStruct(v.Metadata); err == nil {
			member.Metadata = metadataStruct
		}
	}

	return member
}

// MemberListROToProto 批量转换 MemberRO 到 proto Member
func MemberListROToProto(results []memberAppResult.MemberRO) []*memberpb.Member {
	members := make([]*memberpb.Member, len(results))
	for i, v := range results {
		members[i] = MemberROToProto(&v)
	}
	return members
}

// memberStatusToProto 将 domain MemberStatus 转换为 proto TenantsMemberStatus
func memberStatusToProto(status memberDomain.MemberStatus) memberpb.TenantsMemberStatus {
	switch status {
	case memberDomain.MemberStatusInvited:
		return memberpb.TenantsMemberStatus_TENANTS_MEMBER_STATUS_INVITED
	case memberDomain.MemberStatusActive:
		return memberpb.TenantsMemberStatus_TENANTS_MEMBER_STATUS_ACTIVE
	case memberDomain.MemberStatusSuspended:
		return memberpb.TenantsMemberStatus_TENANTS_MEMBER_STATUS_SUSPENDED
	case memberDomain.MemberStatusRemoved:
		return memberpb.TenantsMemberStatus_TENANTS_MEMBER_STATUS_REMOVED
	default:
		return memberpb.TenantsMemberStatus_TENANTS_MEMBER_STATUS_UNSPECIFIED
	}
}

// memberSourceToProto 将 domain MemberSource 转换为 proto TenantsMemberSource
func memberSourceToProto(source memberDomain.MemberSource) memberpb.TenantsMemberSource {
	switch source {
	case memberDomain.MemberSourceManual:
		return memberpb.TenantsMemberSource_TENANTS_MEMBER_SOURCE_MANUAL
	case memberDomain.MemberSourceInvite:
		return memberpb.TenantsMemberSource_TENANTS_MEMBER_SOURCE_INVITE
	case memberDomain.MemberSourceSCIM:
		return memberpb.TenantsMemberSource_TENANTS_MEMBER_SOURCE_SCIM
	case memberDomain.MemberSourceSSO:
		return memberpb.TenantsMemberSource_TENANTS_MEMBER_SOURCE_SSO
	case memberDomain.MemberSourceHRSync:
		return memberpb.TenantsMemberSource_TENANTS_MEMBER_SOURCE_HR_SYNC
	case memberDomain.MemberSourceImport:
		return memberpb.TenantsMemberSource_TENANTS_MEMBER_SOURCE_IMPORT
	default:
		return memberpb.TenantsMemberSource_TENANTS_MEMBER_SOURCE_UNSPECIFIED
	}
}
