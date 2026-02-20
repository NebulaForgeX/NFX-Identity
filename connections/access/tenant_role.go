package access

import (
	"context"
	"fmt"
	tenantrolepb "nfxid/protos/gen/access/tenant_role"
)

type TenantRoleClient struct {
	client tenantrolepb.TenantRoleServiceClient
}

func NewTenantRoleClient(client tenantrolepb.TenantRoleServiceClient) *TenantRoleClient {
	return &TenantRoleClient{client: client}
}

func (c *TenantRoleClient) GetTenantRoleByID(ctx context.Context, id string) (*tenantrolepb.TenantRole, error) {
	resp, err := c.client.GetTenantRoleByID(ctx, &tenantrolepb.GetTenantRoleByIDRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.TenantRole, nil
}

func (c *TenantRoleClient) GetTenantRoleByTenantIDAndRoleKey(ctx context.Context, tenantID, roleKey string) (*tenantrolepb.TenantRole, error) {
	resp, err := c.client.GetTenantRoleByTenantIDAndRoleKey(ctx, &tenantrolepb.GetTenantRoleByTenantIDAndRoleKeyRequest{TenantId: tenantID, RoleKey: roleKey})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.TenantRole, nil
}

func (c *TenantRoleClient) ListTenantRolesByTenantID(ctx context.Context, tenantID string) ([]*tenantrolepb.TenantRole, error) {
	resp, err := c.client.ListTenantRolesByTenantID(ctx, &tenantrolepb.ListTenantRolesByTenantIDRequest{TenantId: tenantID})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.TenantRoles, nil
}

func (c *TenantRoleClient) BatchGetTenantRoles(ctx context.Context, ids []string) ([]*tenantrolepb.TenantRole, error) {
	resp, err := c.client.BatchGetTenantRoles(ctx, &tenantrolepb.BatchGetTenantRolesRequest{Ids: ids})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.TenantRoles, nil
}
