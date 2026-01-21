package directory

import (
	"context"
	"fmt"

	userbadgepb "nfxid/protos/gen/directory/user_badge"
)

// UserBadgeClient UserBadge 客户端
type UserBadgeClient struct {
	client userbadgepb.UserBadgeServiceClient
}

// NewUserBadgeClient 创建 UserBadge 客户端
func NewUserBadgeClient(client userbadgepb.UserBadgeServiceClient) *UserBadgeClient {
	return &UserBadgeClient{client: client}
}

// GetUserBadgeByID 根据ID获取用户徽章
func (c *UserBadgeClient) GetUserBadgeByID(ctx context.Context, id string) (*userbadgepb.UserBadge, error) {
	req := &userbadgepb.GetUserBadgeByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetUserBadgeByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserBadge, nil
}

// GetUserBadgesByUserID 根据用户ID获取用户徽章列表
func (c *UserBadgeClient) GetUserBadgesByUserID(ctx context.Context, userID string) ([]*userbadgepb.UserBadge, error) {
	req := &userbadgepb.GetUserBadgesByUserIDRequest{
		UserId: userID,
	}

	resp, err := c.client.GetUserBadgesByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserBadges, nil
}

// GetUserBadgesByBadgeID 根据徽章ID获取用户徽章列表
func (c *UserBadgeClient) GetUserBadgesByBadgeID(ctx context.Context, badgeID string) ([]*userbadgepb.UserBadge, error) {
	req := &userbadgepb.GetUserBadgesByBadgeIDRequest{
		BadgeId: badgeID,
	}

	resp, err := c.client.GetUserBadgesByBadgeID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserBadges, nil
}