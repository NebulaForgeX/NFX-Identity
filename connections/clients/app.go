package clients

import (
	"context"
	"fmt"

	apppb "nfxid/protos/gen/clients/app"
)

// AppClient App 客户端
type AppClient struct {
	client apppb.AppServiceClient
}

// NewAppClient 创建 App 客户端
func NewAppClient(client apppb.AppServiceClient) *AppClient {
	return &AppClient{client: client}
}

// GetAppByID 根据ID获取应用
func (c *AppClient) GetAppByID(ctx context.Context, id string) (*apppb.App, error) {
	req := &apppb.GetAppByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetAppByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.App, nil
}

// GetAppByAppID 根据应用标识符获取应用
func (c *AppClient) GetAppByAppID(ctx context.Context, appID string) (*apppb.App, error) {
	req := &apppb.GetAppByAppIDRequest{
		AppId: appID,
	}

	resp, err := c.client.GetAppByAppID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.App, nil
}

// GetAppsByTenantID 根据租户ID获取应用列表
func (c *AppClient) GetAppsByTenantID(ctx context.Context, tenantID string) ([]*apppb.App, error) {
	req := &apppb.GetAppsByTenantIDRequest{
		TenantId: tenantID,
	}

	resp, err := c.client.GetAppsByTenantID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Apps, nil
}

// BatchGetApps 批量获取应用
func (c *AppClient) BatchGetApps(ctx context.Context, ids []string) ([]*apppb.App, error) {
	req := &apppb.BatchGetAppsRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetApps(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Apps, nil
}