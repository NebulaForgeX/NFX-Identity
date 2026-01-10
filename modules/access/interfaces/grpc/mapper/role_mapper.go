package mapper

import (
	roleAppResult "nfxid/modules/access/application/roles/results"
	roleDomain "nfxid/modules/access/domain/roles"
	rolepb "nfxid/protos/gen/access/role"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// RoleROToProto 将 RoleRO 转换为 proto Role 消息
func RoleROToProto(v *roleAppResult.RoleRO) *rolepb.Role {
	if v == nil {
		return nil
	}

	role := &rolepb.Role{
		Id:        v.ID.String(),
		Key:       v.Key,
		Name:      v.Name,
		ScopeType: scopeTypeToProto(v.ScopeType),
		IsSystem:  v.IsSystem,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Description != nil {
		role.Description = v.Description
	}

	if v.DeletedAt != nil {
		role.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return role
}

// RoleListROToProto 批量转换 RoleRO 到 proto Role
func RoleListROToProto(results []roleAppResult.RoleRO) []*rolepb.Role {
	roles := make([]*rolepb.Role, len(results))
	for i, v := range results {
		roles[i] = RoleROToProto(&v)
	}
	return roles
}

// scopeTypeToProto 将 domain ScopeType 转换为 proto AccessScopeType
func scopeTypeToProto(scopeType roleDomain.ScopeType) rolepb.AccessScopeType {
	switch scopeType {
	case roleDomain.ScopeTypeTenant:
		return rolepb.AccessScopeType_ACCESS_SCOPE_TYPE_TENANT
	case roleDomain.ScopeTypeApp:
		return rolepb.AccessScopeType_ACCESS_SCOPE_TYPE_APP
	case roleDomain.ScopeTypeGlobal:
		return rolepb.AccessScopeType_ACCESS_SCOPE_TYPE_GLOBAL
	default:
		return rolepb.AccessScopeType_ACCESS_SCOPE_TYPE_UNSPECIFIED
	}
}
