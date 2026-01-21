package clients

import (
	"context"
	"fmt"

	ratelimitpb "nfxid/protos/gen/clients/rate_limit"
)

// RateLimitClient RateLimit 客户端
type RateLimitClient struct {
	client ratelimitpb.RateLimitServiceClient
}

// NewRateLimitClient 创建 RateLimit 客户端
func NewRateLimitClient(client ratelimitpb.RateLimitServiceClient) *RateLimitClient {
	return &RateLimitClient{client: client}
}

// GetRateLimitByID 根据ID获取速率限制
func (c *RateLimitClient) GetRateLimitByID(ctx context.Context, id string) (*ratelimitpb.RateLimit, error) {
	req := &ratelimitpb.GetRateLimitByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetRateLimitByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.RateLimit, nil
}

// GetRateLimitsByAppID 根据应用ID获取速率限制列表
func (c *RateLimitClient) GetRateLimitsByAppID(ctx context.Context, appID string, status *string) ([]*ratelimitpb.RateLimit, error) {
	req := &ratelimitpb.GetRateLimitsByAppIDRequest{
		AppId: appID,
		Status: status,
	}

	resp, err := c.client.GetRateLimitsByAppID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.RateLimits, nil
}