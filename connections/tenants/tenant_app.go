package tenants

import (
	"context"
	"fmt"

	tenantapppb "nfxid/protos/gen/tenants/tenant_app"
)

// TenantAppClient TenantApp 客户端
type TenantAppClient struct {
	client tenantapppb.TenantAppServiceClient
}

// NewTenantAppClient 创建 TenantApp 客户端
func NewTenantAppClient(client tenantapppb.TenantAppServiceClient) *TenantAppClient {
	return &TenantAppClient{client: client}
}

// GetTenantAppByID 根据ID获取租户应用
func (c *TenantAppClient) GetTenantAppByID(ctx context.Context, id string) (*tenantapppb.TenantApp, error) {
	req := &tenantapppb.GetTenantAppByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetTenantAppByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.TenantApp, nil
}

// GetTenantAppsByTenantID 根据租户ID获取租户应用列表
func (c *TenantAppClient) GetTenantAppsByTenantID(ctx context.Context, tenantID string) ([]*tenantapppb.TenantApp, error) {
	req := &tenantapppb.GetTenantAppsByTenantIDRequest{
		TenantId: tenantID,
	}

	resp, err := c.client.GetTenantAppsByTenantID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.TenantApps, nil
}

// GetTenantAppsByAppID 根据应用ID获取租户应用列表
func (c *TenantAppClient) GetTenantAppsByAppID(ctx context.Context, appID string) ([]*tenantapppb.TenantApp, error) {
	req := &tenantapppb.GetTenantAppsByAppIDRequest{
		AppId: appID,
	}

	resp, err := c.client.GetTenantAppsByAppID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.TenantApps, nil
}