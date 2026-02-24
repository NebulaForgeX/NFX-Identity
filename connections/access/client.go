package access

import (
	applicationrolepb "nfxid/protos/gen/access/application_role"
	applicationroleassignmentpb "nfxid/protos/gen/access/application_role_assignment"
	superadminpb "nfxid/protos/gen/access/super_admin"
	tenantrolepb "nfxid/protos/gen/access/tenant_role"
	tenantroleassignmentpb "nfxid/protos/gen/access/tenant_role_assignment"
)

// Client Access 服务 gRPC 客户端聚合
type Client struct {
	SuperAdmin                *SuperAdminClient
	TenantRole                *TenantRoleClient
	TenantRoleAssignment      *TenantRoleAssignmentClient
	ApplicationRole           *ApplicationRoleClient
	ApplicationRoleAssignment *ApplicationRoleAssignmentClient
}

// NewClient 创建 Access 客户端
func NewClient(
	superAdminClient superadminpb.SuperAdminServiceClient,
	tenantRoleClient tenantrolepb.TenantRoleServiceClient,
	tenantRoleAssignmentClient tenantroleassignmentpb.TenantRoleAssignmentServiceClient,
	applicationRoleClient applicationrolepb.ApplicationRoleServiceClient,
	applicationRoleAssignmentClient applicationroleassignmentpb.ApplicationRoleAssignmentServiceClient,
) *Client {
	return &Client{
		SuperAdmin:                NewSuperAdminClient(superAdminClient),
		TenantRole:                NewTenantRoleClient(tenantRoleClient),
		TenantRoleAssignment:      NewTenantRoleAssignmentClient(tenantRoleAssignmentClient),
		ApplicationRole:           NewApplicationRoleClient(applicationRoleClient),
		ApplicationRoleAssignment: NewApplicationRoleAssignmentClient(applicationRoleAssignmentClient),
	}
}
