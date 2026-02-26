package access

import (
	applicationrolepb "nfxidentity/protos/gen/access/application_role"
	applicationroleassignmentpb "nfxidentity/protos/gen/access/application_role_assignment"
	superadminpb "nfxidentity/protos/gen/access/super_admin"
	tenantrolepb "nfxidentity/protos/gen/access/tenant_role"
	tenantroleassignmentpb "nfxidentity/protos/gen/access/tenant_role_assignment"
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
