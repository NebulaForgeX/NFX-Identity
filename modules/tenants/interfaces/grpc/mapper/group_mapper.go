package mapper

import (
	groupAppResult "nfxid/modules/tenants/application/groups/results"
	groupDomain "nfxid/modules/tenants/domain/groups"
	grouppb "nfxid/protos/gen/tenants/group"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GroupROToProto 将 GroupRO 转换为 proto Group 消息
func GroupROToProto(v *groupAppResult.GroupRO) *grouppb.Group {
	if v == nil {
		return nil
	}

	group := &grouppb.Group{
		Id:        v.ID.String(),
		GroupId:   v.GroupID,
		TenantId:  v.TenantID.String(),
		Name:      v.Name,
		Type:      groupTypeToProto(v.Type),
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.ParentGroupID != nil {
		parentID := v.ParentGroupID.String()
		group.ParentGroupId = &parentID
	}

	if v.Description != nil {
		group.Description = v.Description
	}

	if v.CreatedBy != nil {
		createdBy := v.CreatedBy.String()
		group.CreatedBy = &createdBy
	}

	if v.DeletedAt != nil {
		group.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	if v.Metadata != nil && len(v.Metadata) > 0 {
		if metadataStruct, err := structpb.NewStruct(v.Metadata); err == nil {
			group.Metadata = metadataStruct
		}
	}

	return group
}

// GroupListROToProto 批量转换 GroupRO 到 proto Group
func GroupListROToProto(results []groupAppResult.GroupRO) []*grouppb.Group {
	groups := make([]*grouppb.Group, len(results))
	for i, v := range results {
		groups[i] = GroupROToProto(&v)
	}
	return groups
}

// groupTypeToProto 将 domain GroupType 转换为 proto TenantsGroupType
func groupTypeToProto(groupType groupDomain.GroupType) grouppb.TenantsGroupType {
	switch groupType {
	case groupDomain.GroupTypeDepartment:
		return grouppb.TenantsGroupType_TENANTS_GROUP_TYPE_DEPARTMENT
	case groupDomain.GroupTypeTeam:
		return grouppb.TenantsGroupType_TENANTS_GROUP_TYPE_TEAM
	case groupDomain.GroupTypeGroup:
		return grouppb.TenantsGroupType_TENANTS_GROUP_TYPE_GROUP
	case groupDomain.GroupTypeOther:
		return grouppb.TenantsGroupType_TENANTS_GROUP_TYPE_OTHER
	default:
		return grouppb.TenantsGroupType_TENANTS_GROUP_TYPE_UNSPECIFIED
	}
}
