package directory

import (
	"context"
	"fmt"

	badgepb "nfxid/protos/gen/directory/badge"
)

// BadgeClient Badge 客户端
type BadgeClient struct {
	client badgepb.BadgeServiceClient
}

// NewBadgeClient 创建 Badge 客户端
func NewBadgeClient(client badgepb.BadgeServiceClient) *BadgeClient {
	return &BadgeClient{client: client}
}

// GetBadgeByID 根据ID获取徽章
func (c *BadgeClient) GetBadgeByID(ctx context.Context, id string) (*badgepb.Badge, error) {
	req := &badgepb.GetBadgeByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetBadgeByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Badge, nil
}

// GetBadgeByName 根据名称获取徽章
func (c *BadgeClient) GetBadgeByName(ctx context.Context, name string) (*badgepb.Badge, error) {
	req := &badgepb.GetBadgeByNameRequest{
		Name: name,
	}

	resp, err := c.client.GetBadgeByName(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Badge, nil
}

// GetAllBadges 获取所有徽章列表
func (c *BadgeClient) GetAllBadges(ctx context.Context, category *string, isSystem *bool) ([]*badgepb.Badge, error) {
	req := &badgepb.GetAllBadgesRequest{}

	if category != nil {
		req.Category = category
	}
	if isSystem != nil {
		req.IsSystem = isSystem
	}

	resp, err := c.client.GetAllBadges(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Badges, nil
}

// BatchGetBadges 批量获取徽章
func (c *BadgeClient) BatchGetBadges(ctx context.Context, ids []string) ([]*badgepb.Badge, error) {
	req := &badgepb.BatchGetBadgesRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetBadges(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Badges, nil
}