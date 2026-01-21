package tenants

import (
	"context"
	"fmt"

	tenantpb "nfxid/protos/gen/tenants/tenant"
)

// TenantClient Tenant 客户端
type TenantClient struct {
	client tenantpb.TenantServiceClient
}

// NewTenantClient 创建 Tenant 客户端
func NewTenantClient(client tenantpb.TenantServiceClient) *TenantClient {
	return &TenantClient{client: client}
}

// GetTenantByID 根据ID获取租户
func (c *TenantClient) GetTenantByID(ctx context.Context, id string) (*tenantpb.Tenant, error) {
	req := &tenantpb.GetTenantByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetTenantByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Tenant, nil
}

// GetTenantByTenantID 根据租户标识符获取租户
func (c *TenantClient) GetTenantByTenantID(ctx context.Context, tenantID string) (*tenantpb.Tenant, error) {
	req := &tenantpb.GetTenantByTenantIDRequest{
		TenantId: tenantID,
	}

	resp, err := c.client.GetTenantByTenantID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Tenant, nil
}

// BatchGetTenants 批量获取租户
func (c *TenantClient) BatchGetTenants(ctx context.Context, ids []string) ([]*tenantpb.Tenant, error) {
	req := &tenantpb.BatchGetTenantsRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetTenants(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Tenants, nil
}