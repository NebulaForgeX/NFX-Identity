package mapper

import (
	rolePermissionAppResult "nfxid/modules/access/application/role_permissions/results"
	rolepermissionpb "nfxid/protos/gen/access/role_permission"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// RolePermissionROToProto 将 RolePermissionRO 转换为 proto RolePermission 消息
func RolePermissionROToProto(v *rolePermissionAppResult.RolePermissionRO) *rolepermissionpb.RolePermission {
	if v == nil {
		return nil
	}

	rolePermission := &rolepermissionpb.RolePermission{
		Id:           v.ID.String(),
		RoleId:       v.RoleID.String(),
		PermissionId: v.PermissionID.String(),
		CreatedAt:    timestamppb.New(v.CreatedAt),
	}

	if v.CreatedBy != nil {
		createdByStr := v.CreatedBy.String()
		rolePermission.CreatedBy = &createdByStr
	}

	return rolePermission
}

// RolePermissionListROToProto 批量转换 RolePermissionRO 到 proto RolePermission
func RolePermissionListROToProto(results []rolePermissionAppResult.RolePermissionRO) []*rolepermissionpb.RolePermission {
	rolePermissions := make([]*rolepermissionpb.RolePermission, len(results))
	for i, v := range results {
		rolePermissions[i] = RolePermissionROToProto(&v)
	}
	return rolePermissions
}
