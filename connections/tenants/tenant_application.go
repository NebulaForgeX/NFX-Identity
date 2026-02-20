package tenants

import (
	"context"
	"fmt"
	tenantapplicationpb "nfxid/protos/gen/tenants/tenant_application"
)

type TenantApplicationClient struct {
	client tenantapplicationpb.TenantApplicationServiceClient
}

func NewTenantApplicationClient(client tenantapplicationpb.TenantApplicationServiceClient) *TenantApplicationClient {
	return &TenantApplicationClient{client: client}
}

func (c *TenantApplicationClient) GetTenantApplicationByID(ctx context.Context, id string) (*tenantapplicationpb.TenantApplication, error) {
	resp, err := c.client.GetTenantApplicationByID(ctx, &tenantapplicationpb.GetTenantApplicationByIDRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.TenantApplication, nil
}

func (c *TenantApplicationClient) GetTenantApplicationsByTenantID(ctx context.Context, tenantID string) ([]*tenantapplicationpb.TenantApplication, error) {
	resp, err := c.client.GetTenantApplicationsByTenantID(ctx, &tenantapplicationpb.GetTenantApplicationsByTenantIDRequest{TenantId: tenantID})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.TenantApplications, nil
}

func (c *TenantApplicationClient) GetTenantApplicationsByApplicationID(ctx context.Context, applicationID string) ([]*tenantapplicationpb.TenantApplication, error) {
	resp, err := c.client.GetTenantApplicationsByApplicationID(ctx, &tenantapplicationpb.GetTenantApplicationsByApplicationIDRequest{ApplicationId: applicationID})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.TenantApplications, nil
}
