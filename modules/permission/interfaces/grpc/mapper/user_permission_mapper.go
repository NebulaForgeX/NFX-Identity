package mapper

import (
	"nfxid/enums"
	userPermissionAppViews "nfxid/modules/permission/application/user_permission/views"
	permissionpb "nfxid/protos/gen/permission/permission"
	userpermissionpb "nfxid/protos/gen/permission/user_permission"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserPermissionViewToProto 将 UserPermissionView 转换为 proto UserPermission 消息
func UserPermissionViewToProto(v *userPermissionAppViews.UserPermissionView, includePermission bool) *userpermissionpb.UserPermission {
	if v == nil {
		return nil
	}

	up := &userpermissionpb.UserPermission{
		Id:           v.ID.String(),
		UserId:       v.UserID.String(),
		PermissionId: v.PermissionID.String(),
		CreatedAt:    timestamppb.New(v.CreatedAt),
	}

	if includePermission && v.Tag != "" {
		// 如果包含权限信息，创建嵌套的 Permission 消息
		up.Permission = &permissionpb.Permission{
			Id:   v.PermissionID.String(),
			Tag:  v.Tag,
			Name: v.Name,
		}
		var zeroCategory enums.PermissionCategory
		if v.Category != zeroCategory {
			categoryStr := string(v.Category)
			up.Permission.Category = &categoryStr
		}
	}

	return up
}
