package system

import (
	"context"
	"fmt"

	systemstatepb "nfxid/protos/gen/system/system_state"
)

// SystemStateClient SystemState 客户端
type SystemStateClient struct {
	client systemstatepb.SystemStateServiceClient
}

// NewSystemStateClient 创建 SystemState 客户端
func NewSystemStateClient(client systemstatepb.SystemStateServiceClient) *SystemStateClient {
	return &SystemStateClient{client: client}
}

// GetSystemStateByID 根据ID获取系统状态
func (c *SystemStateClient) GetSystemStateByID(ctx context.Context, id string) (*systemstatepb.SystemState, error) {
	req := &systemstatepb.GetSystemStateByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetSystemStateByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.SystemState, nil
}

// GetSystemStateByKey 根据键获取系统状态
func (c *SystemStateClient) GetSystemStateByKey(ctx context.Context, key string) (*systemstatepb.SystemState, error) {
	req := &systemstatepb.GetSystemStateByKeyRequest{
		Key: key,
	}

	resp, err := c.client.GetSystemStateByKey(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.SystemState, nil
}

// GetAllSystemStates 获取所有系统状态列表
func (c *SystemStateClient) GetAllSystemStates(ctx context.Context, category *string) ([]*systemstatepb.SystemState, error) {
	req := &systemstatepb.GetAllSystemStatesRequest{
		Category: category,
	}

	resp, err := c.client.GetAllSystemStates(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.SystemStates, nil
}