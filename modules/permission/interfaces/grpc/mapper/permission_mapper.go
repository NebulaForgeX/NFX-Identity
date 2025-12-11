package mapper

import (
	permissionAppViews "nfxid/modules/permission/application/permission/views"
	permissionpb "nfxid/protos/gen/permission/permission"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// PermissionViewToProto 将 PermissionView 转换为 proto Permission 消息
func PermissionViewToProto(v *permissionAppViews.PermissionView) *permissionpb.Permission {
	if v == nil {
		return nil
	}

	permission := &permissionpb.Permission{
		Id:        v.ID.String(),
		Tag:       v.Tag,
		Name:      v.Name,
		IsSystem:  v.IsSystem,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Description != "" {
		permission.Description = &v.Description
	}
	if v.Category != "" {
		permission.Category = &v.Category
	}
	if v.UpdatedAt.IsZero() {
		permission.UpdatedAt = timestamppb.New(v.CreatedAt)
	}

	return permission
}
