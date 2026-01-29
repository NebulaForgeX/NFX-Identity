package access

import (
	actionpb "nfxid/protos/gen/access/action"
	actionrequirementpb "nfxid/protos/gen/access/action_requirement"
	grantpb "nfxid/protos/gen/access/grant"
	permissionpb "nfxid/protos/gen/access/permission"
	rolepb "nfxid/protos/gen/access/role"
	rolepermissionpb "nfxid/protos/gen/access/role_permission"
	scopepb "nfxid/protos/gen/access/scope"
	scopepermissionpb "nfxid/protos/gen/access/scope_permission"
)

// Client Access 服务客户端
type Client struct {
	Action            *ActionClient
	ActionRequirement *ActionRequirementClient
	Role              *RoleClient
	Permission        *PermissionClient
	Grant             *GrantClient
	RolePermission    *RolePermissionClient
	Scope             *ScopeClient
	ScopePermission   *ScopePermissionClient
}

// NewClient 创建 Access 客户端
func NewClient(
	actionClient actionpb.ActionServiceClient,
	actionRequirementClient actionrequirementpb.ActionRequirementServiceClient,
	roleClient rolepb.RoleServiceClient,
	permissionClient permissionpb.PermissionServiceClient,
	grantClient grantpb.GrantServiceClient,
	rolePermissionClient rolepermissionpb.RolePermissionServiceClient,
	scopeClient scopepb.ScopeServiceClient,
	scopePermissionClient scopepermissionpb.ScopePermissionServiceClient,
) *Client {
	return &Client{
		Action:            NewActionClient(actionClient),
		ActionRequirement: NewActionRequirementClient(actionRequirementClient),
		Role:              NewRoleClient(roleClient),
		Permission:        NewPermissionClient(permissionClient),
		Grant:             NewGrantClient(grantClient),
		RolePermission:    NewRolePermissionClient(rolePermissionClient),
		Scope:             NewScopeClient(scopeClient),
		ScopePermission:   NewScopePermissionClient(scopePermissionClient),
	}
}
