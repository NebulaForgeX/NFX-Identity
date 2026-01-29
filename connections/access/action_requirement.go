package access

import (
	"context"
	"fmt"

	actionrequirementpb "nfxid/protos/gen/access/action_requirement"
)

// ActionRequirementClient ActionRequirement 客户端
type ActionRequirementClient struct {
	client actionrequirementpb.ActionRequirementServiceClient
}

// NewActionRequirementClient 创建 ActionRequirement 客户端
func NewActionRequirementClient(client actionrequirementpb.ActionRequirementServiceClient) *ActionRequirementClient {
	return &ActionRequirementClient{client: client}
}

// CreateActionRequirement 创建 ActionRequirement
func (c *ActionRequirementClient) CreateActionRequirement(ctx context.Context, actionID, permissionID string, groupID int32) (string, error) {
	if groupID == 0 {
		groupID = 1
	}
	req := &actionrequirementpb.CreateActionRequirementRequest{
		ActionId:     actionID,
		PermissionId: permissionID,
		GroupId:      groupID,
	}
	resp, err := c.client.CreateActionRequirement(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.ActionRequirement.Id, nil
}

// GetActionRequirementByID 根据 ID 获取 ActionRequirement
func (c *ActionRequirementClient) GetActionRequirementByID(ctx context.Context, id string) (*actionrequirementpb.ActionRequirement, error) {
	req := &actionrequirementpb.GetActionRequirementByIDRequest{Id: id}
	resp, err := c.client.GetActionRequirementByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.ActionRequirement, nil
}

// GetByActionID 根据 ActionID 获取 ActionRequirement 列表
func (c *ActionRequirementClient) GetByActionID(ctx context.Context, actionID string) ([]*actionrequirementpb.ActionRequirement, error) {
	req := &actionrequirementpb.GetByActionIDRequest{ActionId: actionID}
	resp, err := c.client.GetByActionID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.ActionRequirements, nil
}

// GetByPermissionID 根据 PermissionID 获取 ActionRequirement 列表
func (c *ActionRequirementClient) GetByPermissionID(ctx context.Context, permissionID string) ([]*actionrequirementpb.ActionRequirement, error) {
	req := &actionrequirementpb.GetByPermissionIDRequest{PermissionId: permissionID}
	resp, err := c.client.GetByPermissionID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.ActionRequirements, nil
}

// DeleteActionRequirement 删除 ActionRequirement
func (c *ActionRequirementClient) DeleteActionRequirement(ctx context.Context, id string) error {
	req := &actionrequirementpb.DeleteActionRequirementRequest{Id: id}
	_, err := c.client.DeleteActionRequirement(ctx, req)
	if err != nil {
		return fmt.Errorf("gRPC call failed: %w", err)
	}
	return nil
}

// DeleteByActionIDAndPermissionID 按 ActionID 与 PermissionID 删除
func (c *ActionRequirementClient) DeleteByActionIDAndPermissionID(ctx context.Context, actionID, permissionID string) error {
	req := &actionrequirementpb.DeleteByActionIDAndPermissionIDRequest{
		ActionId:     actionID,
		PermissionId: permissionID,
	}
	_, err := c.client.DeleteByActionIDAndPermissionID(ctx, req)
	if err != nil {
		return fmt.Errorf("gRPC call failed: %w", err)
	}
	return nil
}

// DeleteByActionID 按 ActionID 删除所有关联
func (c *ActionRequirementClient) DeleteByActionID(ctx context.Context, actionID string) error {
	req := &actionrequirementpb.DeleteByActionIDRequest{ActionId: actionID}
	_, err := c.client.DeleteByActionID(ctx, req)
	if err != nil {
		return fmt.Errorf("gRPC call failed: %w", err)
	}
	return nil
}
