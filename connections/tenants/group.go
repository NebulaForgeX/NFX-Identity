package tenants

import (
	"context"
	"fmt"

	grouppb "nfxid/protos/gen/tenants/group"
)

// GroupClient Group 客户端
type GroupClient struct {
	client grouppb.GroupServiceClient
}

// NewGroupClient 创建 Group 客户端
func NewGroupClient(client grouppb.GroupServiceClient) *GroupClient {
	return &GroupClient{client: client}
}

// GetGroupByID 根据ID获取组
func (c *GroupClient) GetGroupByID(ctx context.Context, id string) (*grouppb.Group, error) {
	req := &grouppb.GetGroupByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetGroupByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Group, nil
}

// GetGroupsByTenantID 根据租户ID获取组列表
func (c *GroupClient) GetGroupsByTenantID(ctx context.Context, tenantID string, parentID *string) ([]*grouppb.Group, error) {
	req := &grouppb.GetGroupsByTenantIDRequest{
		TenantId: tenantID,
		ParentId: parentID,
	}

	resp, err := c.client.GetGroupsByTenantID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Groups, nil
}

// BatchGetGroups 批量获取组
func (c *GroupClient) BatchGetGroups(ctx context.Context, ids []string) ([]*grouppb.Group, error) {
	req := &grouppb.BatchGetGroupsRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetGroups(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Groups, nil
}