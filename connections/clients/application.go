package clients

import (
	"context"
	"fmt"
	applicationpb "nfxid/protos/gen/clients/application"
)

// ApplicationClient 应用 gRPC 客户端
type ApplicationClient struct {
	client applicationpb.ApplicationServiceClient
}

// NewApplicationClient 创建 Application 客户端
func NewApplicationClient(client applicationpb.ApplicationServiceClient) *ApplicationClient {
	return &ApplicationClient{client: client}
}

// GetApplicationByID 根据ID获取应用
func (c *ApplicationClient) GetApplicationByID(ctx context.Context, id string) (*applicationpb.Application, error) {
	resp, err := c.client.GetApplicationByID(ctx, &applicationpb.GetApplicationByIDRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.Application, nil
}

// GetApplicationByApplicationID 根据应用标识符获取应用
func (c *ApplicationClient) GetApplicationByApplicationID(ctx context.Context, applicationID string) (*applicationpb.Application, error) {
	resp, err := c.client.GetApplicationByApplicationID(ctx, &applicationpb.GetApplicationByApplicationIDRequest{ApplicationId: applicationID})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.Application, nil
}

// GetApplicationsByTenantID 根据租户ID获取应用列表
func (c *ApplicationClient) GetApplicationsByTenantID(ctx context.Context, tenantID string, env *applicationpb.ClientsEnvironment, status *applicationpb.ClientsAppStatus) ([]*applicationpb.Application, error) {
	req := &applicationpb.GetApplicationsByTenantIDRequest{TenantId: tenantID}
	if env != nil {
		req.Environment = env
	}
	if status != nil {
		req.Status = status
	}
	resp, err := c.client.GetApplicationsByTenantID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.Applications, nil
}

// BatchGetApplications 批量获取应用
func (c *ApplicationClient) BatchGetApplications(ctx context.Context, ids []string) ([]*applicationpb.Application, error) {
	resp, err := c.client.BatchGetApplications(ctx, &applicationpb.BatchGetApplicationsRequest{Ids: ids})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.Applications, nil
}
