package common

import (
	"context"
	"fmt"

	healthpb "nfxid/protos/gen/common/health"
)

// HealthClient 健康检查客户端
type HealthClient struct {
	client healthpb.HealthServiceClient
}

// NewHealthClient 创建健康检查客户端
func NewHealthClient(client healthpb.HealthServiceClient) *HealthClient {
	return &HealthClient{client: client}
}

// GetHealth 获取服务健康状态（包括数据库、Redis等基础设施）
func (c *HealthClient) GetHealth(ctx context.Context) (*healthpb.GetHealthResponse, error) {
	req := &healthpb.GetHealthRequest{}

	resp, err := c.client.GetHealth(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp, nil
}
