package mapper

import (
	scopePermissionAppResult "nfxid/modules/access/application/scope_permissions/results"
	scopepermissionpb "nfxid/protos/gen/access/scope_permission"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// ScopePermissionROToProto 将 ScopePermissionRO 转换为 proto ScopePermission 消息
func ScopePermissionROToProto(v *scopePermissionAppResult.ScopePermissionRO) *scopepermissionpb.ScopePermission {
	if v == nil {
		return nil
	}

	return &scopepermissionpb.ScopePermission{
		Id:           v.ID.String(),
		Scope:        v.Scope,
		PermissionId: v.PermissionID.String(),
		CreatedAt:    timestamppb.New(v.CreatedAt),
	}
}

// ScopePermissionListROToProto 批量转换 ScopePermissionRO 到 proto ScopePermission
func ScopePermissionListROToProto(results []scopePermissionAppResult.ScopePermissionRO) []*scopepermissionpb.ScopePermission {
	scopePermissions := make([]*scopepermissionpb.ScopePermission, len(results))
	for i, v := range results {
		scopePermissions[i] = ScopePermissionROToProto(&v)
	}
	return scopePermissions
}
