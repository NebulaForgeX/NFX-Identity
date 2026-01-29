package access

import (
	"context"
	"fmt"

	actionpb "nfxid/protos/gen/access/action"
)

// ActionClient Action 客户端
type ActionClient struct {
	client actionpb.ActionServiceClient
}

// NewActionClient 创建 Action 客户端
func NewActionClient(client actionpb.ActionServiceClient) *ActionClient {
	return &ActionClient{client: client}
}

// CreateAction 创建 Action
func (c *ActionClient) CreateAction(ctx context.Context, key, service, status, name string, description *string, isSystem bool) (string, error) {
	if status == "" {
		status = "active"
	}
	req := &actionpb.CreateActionRequest{
		Key:         key,
		Service:     service,
		Status:      status,
		Name:        name,
		Description: description,
		IsSystem:    isSystem,
	}
	resp, err := c.client.CreateAction(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.Action.Id, nil
}

// GetActionByID 根据 ID 获取 Action
func (c *ActionClient) GetActionByID(ctx context.Context, id string) (*actionpb.Action, error) {
	req := &actionpb.GetActionByIDRequest{Id: id}
	resp, err := c.client.GetActionByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.Action, nil
}

// GetActionByKey 根据 Key 获取 Action
func (c *ActionClient) GetActionByKey(ctx context.Context, key string) (*actionpb.Action, error) {
	req := &actionpb.GetActionByKeyRequest{Key: key}
	resp, err := c.client.GetActionByKey(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.Action, nil
}

// UpdateAction 更新 Action
func (c *ActionClient) UpdateAction(ctx context.Context, id, key, service, status, name string, description *string) (*actionpb.Action, error) {
	req := &actionpb.UpdateActionRequest{
		Id:          id,
		Key:         key,
		Service:     service,
		Status:      status,
		Name:        name,
		Description: description,
	}
	resp, err := c.client.UpdateAction(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.Action, nil
}

// DeleteAction 删除 Action
func (c *ActionClient) DeleteAction(ctx context.Context, id string) error {
	req := &actionpb.DeleteActionRequest{Id: id}
	_, err := c.client.DeleteAction(ctx, req)
	if err != nil {
		return fmt.Errorf("gRPC call failed: %w", err)
	}
	return nil
}
