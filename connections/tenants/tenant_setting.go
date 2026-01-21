package tenants

import (
	"context"
	"fmt"

	tenantsettingpb "nfxid/protos/gen/tenants/tenant_setting"
)

// TenantSettingClient TenantSetting 客户端
type TenantSettingClient struct {
	client tenantsettingpb.TenantSettingServiceClient
}

// NewTenantSettingClient 创建 TenantSetting 客户端
func NewTenantSettingClient(client tenantsettingpb.TenantSettingServiceClient) *TenantSettingClient {
	return &TenantSettingClient{client: client}
}

// GetTenantSettingByID 根据ID获取租户设置
func (c *TenantSettingClient) GetTenantSettingByID(ctx context.Context, id string) (*tenantsettingpb.TenantSetting, error) {
	req := &tenantsettingpb.GetTenantSettingByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetTenantSettingByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.TenantSetting, nil
}

// GetTenantSettingByTenantID 根据租户ID获取租户设置
func (c *TenantSettingClient) GetTenantSettingByTenantID(ctx context.Context, tenantID string) (*tenantsettingpb.TenantSetting, error) {
	req := &tenantsettingpb.GetTenantSettingByTenantIDRequest{
		TenantId: tenantID,
	}

	resp, err := c.client.GetTenantSettingByTenantID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.TenantSetting, nil
}