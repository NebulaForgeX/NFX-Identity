package mapper

import (
	permissionAppResult "nfxid/modules/access/application/permissions/results"
	permissionpb "nfxid/protos/gen/access/permission"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// PermissionROToProto 将 PermissionRO 转换为 proto Permission 消息
func PermissionROToProto(v *permissionAppResult.PermissionRO) *permissionpb.Permission {
	if v == nil {
		return nil
	}

	permission := &permissionpb.Permission{
		Id:        v.ID.String(),
		Key:       v.Key,
		Name:      v.Name,
		IsSystem:  v.IsSystem,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Description != nil {
		permission.Description = v.Description
	}

	if v.DeletedAt != nil {
		permission.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return permission
}

// PermissionListROToProto 批量转换 PermissionRO 到 proto Permission
func PermissionListROToProto(results []permissionAppResult.PermissionRO) []*permissionpb.Permission {
	permissions := make([]*permissionpb.Permission, len(results))
	for i, v := range results {
		permissions[i] = PermissionROToProto(&v)
	}
	return permissions
}
